package postgresrepo

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/pkg/org"
)

var activityUnitsTable = newTable(
	"activity_units",
	[]string{
		"activity_id",
		"commander_id",
		"superior_unit_id",
		"unit_name",
	},
	"",
)

func mapRowToActivityUnit(row scannable) (org.ActivityUnit, uuid.UUID, error) {
	var id uuid.UUID
	var activityID uuid.UUID
	var commanderID uuid.UUID
	var superiorUnitID uuid.UUID
	var unitName string

	err := row.Scan(
		&id,
		&activityID,
		&commanderID,
		&superiorUnitID,
		&unitName,
	)

	if err != nil {
		return org.ActivityUnit{}, uuid.Nil, errors.WithStack(err)
	}

	return org.NewActivityUnit(
		id,
		unitName,
	), superiorUnitID, nil
}

func (pr *PostgresRepository) insertActivityUnit(tx *sql.Tx, activityId, superiorUnitId uuid.UUID, activityUnit org.ActivityUnit) (org.ActivityUnit, error) {
	stmt, err := upsertStatement(tx, activityUnitsTable)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	var row *sql.Row
	if superiorUnitId == uuid.Nil {
		row = stmt.QueryRow(
			activityUnit.ID(),
			activityId,
			uuid.Nil,
			nil,
			activityUnit.UnitName(),
		)
	} else {
		row = stmt.QueryRow(
			activityUnit.ID(),
			activityId,
			uuid.Nil,
			superiorUnitId,
			activityUnit.UnitName(),
		)
	}

	insertedActivityUnit, _, err := mapRowToActivityUnit(row)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return insertedActivityUnit, nil
}

func (pr *PostgresRepository) bulkInsertActivityUnits(tx *sql.Tx, activityId, superiorUnitId uuid.UUID, activityUnits []org.ActivityUnit) ([]org.ActivityUnit, error) {
	stmt, err := bulkUpsertStatement(tx, activityUnitsTable, len(activityUnits))
	if err != nil {
		return nil, errors.WithStack(processError("activityUnit", "", err))
	}

	var values []any
	for _, unit := range activityUnits {
		values = append(values, unit.ID())
		values = append(values, activityId)
		values = append(values, nil)
		if superiorUnitId == uuid.Nil {
			values = append(values, nil)
		} else {
			values = append(values, superiorUnitId)
		}
		values = append(values, unit.UnitName())
	}

	rows, err := stmt.Query(values...)
	if err != nil {
		return nil, errors.WithStack(processError("activityUnit", "", err))
	}

	var insertedActivityUnits []org.ActivityUnit
	for rows.Next() {
		activityUnit, _, err := mapRowToActivityUnit(rows)
		if err != nil {
			return nil, errors.WithStack(processError("activityUnit", "", err))
		}

		insertedActivityUnits = append(insertedActivityUnits, activityUnit)
	}

	return insertedActivityUnits, nil
}

func (pr *PostgresRepository) insertHierarchy(tx *sql.Tx, activityId uuid.UUID, unit org.ActivityUnit) (org.ActivityUnit, error) {
	subordinateUnits := unit.SubordinateUnits()
	if len(subordinateUnits) == 0 {
		return unit, nil
	}

	insertedUnit := org.NewActivityUnit(unit.ID(), unit.UnitName())

	insertedSubUnits, err := pr.bulkInsertActivityUnits(tx, activityId, unit.ID(), subordinateUnits)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	for i, u := range insertedSubUnits {
		for _, s := range subordinateUnits[i].SubordinateUnits() {
			u.AddSubordinateUnit(s)
		}

		iu, err := pr.insertHierarchy(tx, activityId, u)
		if err != nil {
			return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
		}
		insertedUnit.AddSubordinateUnit(iu)
	}

	return insertedUnit, nil
}

func (pr *PostgresRepository) InsertActivityUnit(activityKey string, superiorUnitId uuid.UUID, activityUnit org.ActivityUnit) (org.ActivityUnit, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	activityID, err := pr.idByKey(activitiesTable, activityKey)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	insertedActivityUnit, err := pr.insertActivityUnit(tx, activityID, superiorUnitId, activityUnit)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	for _, u := range activityUnit.SubordinateUnits() {
		insertedActivityUnit.AddSubordinateUnit(u)
	}

	insertedActivityUnit, err = pr.insertHierarchy(tx, activityID, insertedActivityUnit)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	err = tx.Commit()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return insertedActivityUnit, nil
}

func (pr *PostgresRepository) buildHierarchy(node org.ActivityUnit,
	units map[uuid.UUID][]org.ActivityUnit) org.ActivityUnit {

	subunits, ok := units[node.ID()]
	if !ok {
		return node
	}

	for _, u := range subunits {
		u = pr.buildHierarchy(u, units)
		node.AddSubordinateUnit(u)
	}

	return node
}

func (pr *PostgresRepository) SelectActivityHierarchy(activityKey string) (org.ActivityUnit, error) {
	activityID, err := pr.idByKey(activitiesTable, activityKey)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	stmt, err := selectStatementWhere(pr.db, activityUnitsTable, "activity_id")
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	rows, err := stmt.Query(activityID)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	res := make(map[uuid.UUID][]org.ActivityUnit)
	for rows.Next() {
		unit, superiorUnitID, err := mapRowToActivityUnit(rows)
		if err != nil {
			return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
		}
		res[superiorUnitID] = append(res[superiorUnitID], unit)
	}

	r, ok := res[uuid.Nil]
	if !ok {
		return org.ActivityUnit{}, errors.Errorf("activity does not have a root unit: %s", activityID)
	}

	root := pr.buildHierarchy(r[0], res)

	return root, nil
}

func (pr *PostgresRepository) SelectActivityUnit(id uuid.UUID) (org.ActivityUnit, error) {
	stmt, err := selectStatementByID(pr.db, activityUnitsTable)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	row := stmt.QueryRow(id)

	activityUnit, _, err := mapRowToActivityUnit(row)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return activityUnit, nil
}

func (pr *PostgresRepository) UpdateActivityUnit(activityKey string, id, superiorUnitID uuid.UUID, unit org.ActivityUnit) (org.ActivityUnit, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	activityID, err := pr.idByKey(activitiesTable, activityKey)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	stmt, err := updateStatement(tx, activityUnitsTable)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	row := stmt.QueryRow(
		activityID,
		uuid.Nil,
		superiorUnitID,
		unit.UnitName(),
		id,
	)

	insertedActivityUnit, _, err := mapRowToActivityUnit(row)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	err = tx.Commit()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return insertedActivityUnit, nil
}

func (pr *PostgresRepository) DeleteActivityUnit(id uuid.UUID) error {
	tx, err := pr.db.Begin()
	if err != nil {
		return errors.WithStack(processError("activityUnit", "", err))
	}

	stmt, err := deleteStatement(tx, activityUnitsTable)
	if err != nil {
		return errors.WithStack(processError("activityUnit", "", err))
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return errors.WithStack(processError("activityUnit", "", err))
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.WithStack(processError("activityUnit", "", err))
	}

	if affected == 0 {
		return errors.WithStack(
			repositories.ErrNotFound{
				Object: "activityUnit",
				Key:    id.String(),
			},
		)
	}

	err = tx.Commit()
	if err != nil {
		return errors.WithStack(processError("activityUnit", "", err))
	}

	return nil
}
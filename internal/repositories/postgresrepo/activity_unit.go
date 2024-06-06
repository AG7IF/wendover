package postgresrepo

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"

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

func mapRowToActivityUnit(row scannable) (org.ActivityUnit, error) {
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
		return org.ActivityUnit{}, errors.WithStack(err)
	}

	return org.NewActivityUnit(
		id,
		unitName,
	), nil
}

func (pr *PostgresRepository) insertActivityUnit(tx *sql.Tx, activityId, superiorUnitId uuid.UUID, activityUnit org.ActivityUnit) (org.ActivityUnit, error) {
	stmt, err := insertStatement(tx, activityUnitsTable)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	var row *sql.Row
	if superiorUnitId == uuid.Nil {
		row = stmt.QueryRow(
			activityId,
			uuid.Nil,
			nil,
			activityUnit.UnitName(),
		)
	} else {
		row = stmt.QueryRow(
			activityId,
			uuid.Nil,
			superiorUnitId,
			activityUnit.UnitName(),
		)
	}

	insertedActivityUnit, err := mapRowToActivityUnit(row)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return insertedActivityUnit, nil
}

func (pr *PostgresRepository) InsertActivityUnit(activityId, superiorUnitId uuid.UUID, activityUnit org.ActivityUnit) (org.ActivityUnit, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	insertedActivityUnit, err := pr.insertActivityUnit(tx, activityId, superiorUnitId, activityUnit)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	err = tx.Commit()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return insertedActivityUnit, nil
}

func (pr *PostgresRepository) bulkInsertActivityUnits(tx *sql.Tx, activityId, superiorUnitId uuid.UUID, activityUnits []org.ActivityUnit) ([]org.ActivityUnit, error) {
	stmt, err := bulkInsertStatement(tx, activityUnitsTable, len(activityUnits))
	if err != nil {
		return nil, errors.WithStack(processError("activityUnit", "", err))
	}

	var values []any
	for _, unit := range activityUnits {
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
		activityUnit, err := mapRowToActivityUnit(rows)
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

func (pr *PostgresRepository) InsertActivityHierachy(activityId uuid.UUID, rootUnit org.ActivityUnit) (org.ActivityUnit, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	insertedRootUnit, err := pr.insertActivityUnit(tx, activityId, uuid.Nil, rootUnit)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	for _, u := range rootUnit.SubordinateUnits() {
		insertedRootUnit.AddSubordinateUnit(u)
	}

	insertedHierarchy, err := pr.insertHierarchy(tx, activityId, insertedRootUnit)
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	err = tx.Commit()
	if err != nil {
		return org.ActivityUnit{}, errors.WithStack(processError("activityUnit", "", err))
	}

	return insertedHierarchy, nil
}

func (pr *PostgresRepository) SelectActivityHierarchy(activityId uuid.UUID) (org.ActivityUnit, error) {
	//TODO implement me
	panic("implement me")
}

func (pr *PostgresRepository) SelectActivityUnit(id uuid.UUID) (org.ActivityUnit, error) {
	//TODO implement me
	panic("implement me")
}

func (pr *PostgresRepository) UpdateActivityUnit(id uuid.UUID, activity org.ActivityUnit) (org.ActivityUnit, error) {
	//TODO implement me
	panic("implement me")
}

func (pr *PostgresRepository) DeleteActivityUnit(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

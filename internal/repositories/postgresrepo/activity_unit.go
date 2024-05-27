package postgresrepo

import (
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

func (pr *PostgresRepository) InsertActivityUnits(activityId, superiorUnitId uuid.UUID, activityUnits []org.ActivityUnit) ([]org.ActivityUnit, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return nil, errors.WithStack(processError("activityUnit", "", err))
	}

	stmt, err := bulkInsertStatement(tx, activityUnitsTable, len(activityUnits))
	if err != nil {
		return nil, errors.WithStack(processError("activityUnit", "", err))
	}

	var values []any
	for _, unit := range activityUnits {
		values = append(values, unit.unit)
	}
}

func (pr *PostgresRepository) SelectActivityUnitsByActivity(activity org.Activity) ([]org.ActivityUnit, error) {
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

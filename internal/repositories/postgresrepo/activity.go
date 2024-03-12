package postgresrepo

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/pkg/org"
)

var activitiesTable = newTable(
	"activities",
	[]string{
		"key",
		"name",
		"location",
		"start_datetime",
		"end_datetime",
		"cadet_student_fee",
		"cadet_cadre_fee",
		"senior_student_fee",
		"senior_cadre_fee",
	},
	"key",
)

func mapRowToActivity(row scannable) (org.Activity, error) {
	var id uuid.UUID
	var key string
	var name string
	var location string
	var start time.Time
	var end time.Time
	var cadetStudentFee uint
	var cadetCadreFee uint
	var seniorStudentFee uint
	var seniorCadreFee uint

	err := row.Scan(
		&id,
		&key,
		&name,
		&location,
		&start,
		&end,
		&cadetStudentFee,
		&cadetCadreFee,
		&seniorStudentFee,
		&seniorCadreFee,
	)

	if err != nil {
		return org.Activity{}, errors.WithStack(err)
	}

	return org.NewActivity(
		id,
		key,
		name,
		location,
		start,
		end,
		cadetStudentFee,
		cadetCadreFee,
		seniorStudentFee,
		seniorCadreFee,
	), nil
}

func (pr *PostgresRepository) InsertActivity(activity org.Activity) (org.Activity, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", activity.Key(), err))
	}

	stmt, err := insertStatement(tx, activitiesTable)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", activity.Key(), err))
	}

	row := stmt.QueryRow(
		activity.Key(),
		activity.Name(),
		activity.Location(),
		activity.Start(),
		activity.End(),
		activity.CadetStudentFee(),
		activity.CadetCadreFee(),
		activity.SeniorStudentFee(),
		activity.SeniorCadreFee(),
	)

	insertedActivity, err := mapRowToActivity(row)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", activity.Key(), err))
	}

	err = tx.Commit()
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", activity.Key(), err))
	}

	return insertedActivity, nil
}

func (pr *PostgresRepository) SelectActivities() ([]org.Activity, error) {
	stmt, err := selectStatementAll(pr.db, activitiesTable)
	if err != nil {
		return nil, errors.WithStack(processError("activity", "*", err))
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.WithStack(processError("activity", "*", err))
	}

	var activities []org.Activity
	for rows.Next() {
		activity, err := mapRowToActivity(rows)
		if err != nil {
			return nil, errors.WithStack(processError("activity", "*", err))
		}

		activities = append(activities, activity)
	}
	err = rows.Close()
	if err != nil {
		return nil, errors.WithStack(processError("user", "*", err))
	}

	return activities, nil
}

func (pr *PostgresRepository) SelectActivity(key string) (org.Activity, error) {
	stmt, err := selectStatementByKey(pr.db, activitiesTable)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	row := stmt.QueryRow(key)

	activity, err := mapRowToActivity(row)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	return activity, nil
}

func (pr *PostgresRepository) UpdateActivity(key string, activity org.Activity) (org.Activity, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	id, err := pr.idByKey(activitiesTable, key)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	stmt, err := updateStatement(tx, activitiesTable)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	row := stmt.QueryRow(
		activity.Key(),
		activity.Name(),
		activity.Location(),
		activity.Start(),
		activity.End(),
		activity.CadetStudentFee(),
		activity.CadetCadreFee(),
		activity.SeniorStudentFee(),
		activity.SeniorCadreFee(),
		id,
	)

	insertedActivity, err := mapRowToActivity(row)
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	err = tx.Commit()
	if err != nil {
		return org.Activity{}, errors.WithStack(processError("activity", key, err))
	}

	return insertedActivity, nil
}

func (pr *PostgresRepository) DeleteActivity(key string) error {
	tx, err := pr.db.Begin()
	if err != nil {
		return errors.WithStack(processError("activity", key, err))
	}

	stmt, err := deleteStatementByKey(tx, activitiesTable)
	if err != nil {
		return errors.WithStack(processError("activity", key, err))
	}

	res, err := stmt.Exec(key)
	if err != nil {
		return errors.WithStack(processError("activity", key, err))
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.WithStack(processError("activity", key, err))
	}

	if affected == 0 {
		return errors.WithStack(
			repositories.ErrNotFound{
				Object: "activity",
				Key:    key,
			},
		)
	}

	err = tx.Commit()
	if err != nil {
		return errors.WithStack(processError("activity", key, err))
	}

	return nil
}

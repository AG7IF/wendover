package postgresrepo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ag7if/wendover/pkg/auth"
	"github.com/ag7if/wendover/pkg/org"
)

var userRolesTable = newAssocTable(
	"user_roles",
	"user_id",
	"activity_id",
	[]string{
		"role",
	},
)

func selectUserWithRoles(p preparer) (*sql.Stmt, error) {
	query := `
SELECT 
    users.id,
	users.username,
	users.email,
	user_roles.role, 
	activities.id, 
	activities.key, 
	activities.name, 
	activities.location, 
	activities.start_datetime,
	activities.end_datetime,
	activities.cadet_student_fee,
	activities.cadet_cadre_fee,
	activities.senior_student_fee,
	activities.senior_cadre_fee
FROM wendover.users
LEFT JOIN wendover.user_roles ON users.id = user_roles.user_id
LEFT JOIN wendover.activities ON user_roles.activity_id = activities.id
WHERE users.username = $1
`
	stmt, err := p.Prepare(query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stmt, nil
}

func mapRowToUserMapWithRole(row scannable) (map[string]any, auth.UserRole, error) {
	var userID uuid.UUID
	var username string
	var email string
	var role auth.Role
	var activityID uuid.UUID
	var activityKey sql.NullString
	var activityName sql.NullString
	var activityLocation sql.NullString
	var activityStart sql.NullTime
	var activityEnd sql.NullTime
	var activityCadetStudentFee sql.NullInt32
	var activityCadetCadreFee sql.NullInt32
	var activitySeniorStudentFee sql.NullInt32
	var activitySeniorCadreFee sql.NullInt32

	err := row.Scan(
		&userID,
		&username,
		&email,
		&role,
		&activityID,
		&activityKey,
		&activityName,
		&activityLocation,
		&activityStart,
		&activityEnd,
		&activityCadetStudentFee,
		&activityCadetCadreFee,
		&activitySeniorStudentFee,
		&activitySeniorCadreFee,
	)
	if err != nil {
		return nil, auth.UserRole{}, errors.WithStack(err)
	}

	userMap := map[string]any{
		"id":       userID,
		"username": username,
		"email":    email,
	}

	var userRole auth.UserRole

	if role != auth.NilRole {
		activity := org.NewActivity(
			activityID,
			activityKey.String,
			activityName.String,
			activityLocation.String,
			activityStart.Time,
			activityEnd.Time,
			uint(activityCadetStudentFee.Int32),
			uint(activityCadetCadreFee.Int32),
			uint(activitySeniorStudentFee.Int32),
			uint(activitySeniorCadreFee.Int32),
		)

		userRole = auth.NewUserRole(activity, role)
	}

	return userMap, userRole, nil
}

func (pr *PostgresRepository) AddUserRole(username, activityKey string, role auth.Role) (auth.User, error) {
	userID, err := pr.idByKey(usersTable, username)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	activityID, err := pr.idByKey(activitiesTable, activityKey)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	tx, err := pr.db.Begin()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	stmt, err := insertAssocStatement(tx, userRolesTable)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	_, err = stmt.Exec(userID, activityID, role)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	err = tx.Commit()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	return pr.SelectUser(username)
}

func (pr *PostgresRepository) RemoveUserRole(username, activityKey string) (auth.User, error) {
	userID, err := pr.idByKey(usersTable, username)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	activityID, err := pr.idByKey(activitiesTable, activityKey)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	tx, err := pr.db.Begin()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	stmt, err := deleteAssociation(tx, userRolesTable)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	_, err = stmt.Exec(userID, activityID)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	err = tx.Commit()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user_role", fmt.Sprintf("%s|%s", username, activityKey), err))
	}

	return pr.SelectUser(username)
}

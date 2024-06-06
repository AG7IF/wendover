package postgresrepo

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ag7if/wendover/internal/repositories"
	"github.com/ag7if/wendover/pkg/auth"
)

var usersTable = newTable(
	"users",
	[]string{
		"username",
	},
	"username",
)

func mapRowToUserMap(row scannable) (map[string]any, error) {
	var id uuid.UUID
	var username string

	err := row.Scan(
		&id,
		&username,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userMap := map[string]any{
		"id":       id,
		"username": username,
	}

	return userMap, nil
}

func (pr *PostgresRepository) InsertUser(user auth.User) (auth.User, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", user.Username(), err))
	}

	stmt, err := insertStatement(tx, usersTable)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", user.Username(), err))
	}

	row := stmt.QueryRow(user.Username())
	insertedUser, err := mapRowToUserMap(row)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", user.Username(), err))
	}

	err = tx.Commit()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", user.Username(), err))
	}

	return auth.NewUser(
		insertedUser["id"].(uuid.UUID),
		insertedUser["username"].(string),
		nil,
	), nil
}

func (pr *PostgresRepository) SelectUsers() ([]auth.User, error) {
	stmt, err := selectStatementAll(pr.db, usersTable)
	if err != nil {
		return nil, errors.WithStack(processError("user", "*", err))
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.WithStack(processError("user", "*", err))
	}

	var users []auth.User
	for rows.Next() {
		userMap, err := mapRowToUserMap(rows)
		if err != nil {
			return nil, errors.WithStack(processError("user", "*", err))
		}

		user := auth.NewUser(userMap["id"].(uuid.UUID), userMap["username"].(string), nil)
		users = append(users, user)
	}
	err = rows.Close()
	if err != nil {
		return nil, errors.WithStack(processError("user", "*", err))
	}

	return users, nil
}

func (pr *PostgresRepository) SelectUser(username string) (auth.User, error) {
	stmt, err := selectUserWithRoles(pr.db)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	rows, err := stmt.Query(username)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	var userMap map[string]any
	var userRoles map[string]auth.UserRole
	for rows.Next() {
		var userRole auth.UserRole
		if userMap == nil {
			userMap, userRole, err = mapRowToUserMapWithRole(rows)
		} else {
			_, userRole, err = mapRowToUserMapWithRole(rows)

		}
		if err != nil {
			return auth.User{}, errors.WithStack(processError("user", username, err))
		}

		roleActivity := userRole.Activity()

		if roleActivity.ID() == uuid.Nil {
			continue
		}

		if userRoles == nil {
			userRoles = make(map[string]auth.UserRole)
		}

		userRoles[roleActivity.Key()] = userRole
	}

	return auth.NewUser(userMap["id"].(uuid.UUID), userMap["username"].(string), userRoles), nil
}

func (pr *PostgresRepository) UpdateUser(username string, user auth.User) (auth.User, error) {
	id, err := pr.idByKey(usersTable, username)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	tx, err := pr.db.Begin()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	stmt, err := updateStatement(tx, usersTable)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	row := stmt.QueryRow(
		user.Username(),
		id,
	)
	userMap, err := mapRowToUserMap(row)
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	err = tx.Commit()
	if err != nil {
		return auth.User{}, errors.WithStack(processError("user", username, err))
	}

	return pr.SelectUser(userMap["username"].(string))
}

func (pr *PostgresRepository) DeleteUser(username string) error {
	tx, err := pr.db.Begin()
	if err != nil {
		return errors.WithStack(processError("user", username, err))
	}

	stmt, err := deleteStatementByKey(tx, usersTable)
	if err != nil {
		return errors.WithStack(processError("user", username, err))
	}

	res, err := stmt.Exec(username)
	if err != nil {
		return errors.WithStack(processError("user", username, err))
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.WithStack(processError("user", username, err))
	}

	if affected == 0 {
		return errors.WithStack(
			repositories.ErrNotFound{
				Object: "user",
				Key:    username,
			},
		)
	}

	err = tx.Commit()
	if err != nil {
		return errors.WithStack(processError("user", username, err))
	}

	return nil
}

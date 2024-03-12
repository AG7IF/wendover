package postgresrepo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (pr *PostgresRepository) idByKey(t table, key string) (uuid.UUID, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE %s = $1;", t.TableName(), t.KeyColumn())

	stmt, err := pr.db.Prepare(query)
	if err != nil {
		return uuid.Nil, err
	}

	row := stmt.QueryRow(key)
	var id uuid.UUID
	err = row.Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

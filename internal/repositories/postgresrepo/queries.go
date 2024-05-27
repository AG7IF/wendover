package postgresrepo

import (
	"database/sql"
	"fmt"
	"strings"
)

const schemaName = "wendover"

type table struct {
	tableName string
	columns   []string
	keyColumn string
}

func newTable(tableName string, columns []string, keyColumn string) table {
	return table{
		tableName: tableName,
		columns:   columns,
		keyColumn: keyColumn,
	}
}

func newAssocTable(tableName string, fk1Name, fk2Name string, columns []string) table {
	cols := []string{fk1Name, fk2Name}
	cols = append(cols, columns...)
	return table{
		tableName: tableName,
		columns:   cols,
	}
}

func (t table) TableName() string {
	return fmt.Sprintf("%s.%s", schemaName, t.tableName)
}

func (t table) Columns() []string {
	return t.columns
}

func (t table) KeyColumn() string {
	return t.keyColumn
}

type scannable interface {
	Scan(dest ...any) error
}

type preparer interface {
	Prepare(string) (*sql.Stmt, error)
}

func insertStatement(p preparer, t table) (*sql.Stmt, error) {
	var placeholders []string
	for i := range t.Columns() {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id,%s;",
		t.TableName(),
		strings.Join(t.Columns(), ","),
		strings.Join(placeholders, ","),
		strings.Join(t.Columns(), ","),
	)

	return p.Prepare(query)
}

func bulkInsertStatement(p preparer, t table, length int) (*sql.Stmt, error) {

	n := 0
	var values []string
	for i := 0; i < length; i++ {
		var placeholders []string
		for range t.Columns() {
			n++
			placeholders = append(placeholders, fmt.Sprintf("$%d", n))
		}

		v := fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
		values = append(values, v)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s RETURNING, id,%s;",
		t.TableName(),
		strings.Join(t.Columns(), ","),
		strings.Join(values, ","),
		strings.Join(t.Columns(), ","),
	)

	return p.Prepare(query)
}

// insertAssocStatement prepares an INSERT query for an association table. The key differences between this query and
// the query that is prepared by insertStatement are:
//  1. the RETURNING clause is omitted, and
//  2. an ON CONFLICT UPDATE clause is included.
//
// The ON CONFLICT UPDATE clause assumes that the first two columns in the table.columns slice are the pk/fk pair
// and should not be updated. Therefore, SET clauses are only generated for everything after the first two columns
// (if any).
func insertAssocStatement(p preparer, t table) (*sql.Stmt, error) {
	var placeholders []string
	for i := range t.Columns() {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	var setClauses []string
	if len(t.Columns()) > 2 {
		for i, v := range t.Columns()[2:] {
			set := fmt.Sprintf("SET %s = $%d", v, (i+1)+2)
			setClauses = append(setClauses, set)
		}
	}

	var onConflict string
	if setClauses == nil {
		onConflict = "ON CONFLICT DO NOTHING"
	} else {
		onConflict = fmt.Sprintf("ON CONFLICT (%s, %s) DO UPDATE %s",
			t.Columns()[0],
			t.Columns()[1],
			strings.Join(setClauses, ","),
		)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) %s;",
		t.TableName(),
		strings.Join(t.Columns(), ","),
		strings.Join(placeholders, ","),
		onConflict,
	)

	return p.Prepare(query)
}

func selectStatementAll(p preparer, t table) (*sql.Stmt, error) {
	query := fmt.Sprintf("SELECT id,%s FROM %s;",
		strings.Join(t.Columns(), ","),
		t.TableName())

	return p.Prepare(query)
}

func selectStatementByID(p preparer, t table) (*sql.Stmt, error) {
	query := fmt.Sprintf("SELECT id,%s FROM %s WHERE id = $1;",
		strings.Join(t.Columns(), ","),
		t.TableName())

	return p.Prepare(query)
}

func selectStatementByKey(p preparer, t table) (*sql.Stmt, error) {
	query := fmt.Sprintf("SELECT id,%s FROM %s WHERE %s = $1;",
		strings.Join(t.Columns(), ","),
		t.TableName(),
		t.KeyColumn())

	return p.Prepare(query)
}

func selectStatementWhere(p preparer, t table, col string) (*sql.Stmt, error) {
	query := fmt.Sprintf("SELECT id,%s FROM %s WHERE %s = $1",
		strings.Join(t.Columns(), ","),
		t.TableName(),
		col,
	)

	return p.Prepare(query)
}

func selectStatementWhereIn(p preparer, t table, col string) (*sql.Stmt, error) {
	query := fmt.Sprintf("SELECT id,%s FROM %s WHERE %s IN ($1)",
		strings.Join(t.Columns(), ","),
		t.TableName(),
		col,
	)

	return p.Prepare(query)
}

func updateStatement(p preparer, t table) (*sql.Stmt, error) {
	var placeholders []string
	for i, v := range t.Columns() {
		placeholders = append(placeholders, fmt.Sprintf("%s=$%d", v, i+1))
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id,%s;",
		t.TableName(),
		strings.Join(placeholders, ","),
		len(t.Columns())+1,
		strings.Join(t.Columns(), ","),
	)

	return p.Prepare(query)
}

func deleteStatement(p preparer, t table) (*sql.Stmt, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		t.TableName())

	return p.Prepare(query)
}

func deleteStatementByKey(p preparer, t table) (*sql.Stmt, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1",
		t.TableName(),
		t.KeyColumn(),
	)

	return p.Prepare(query)
}

// deleteAssociation will prepare a DELETE statement for association (many-to-many join) tables.
// It assumes that the first two columns in the table.columns slice are the foreign/primary key pairs of such tables.
// Use newAssocTable() to ensure that this slice is properly created in the table object.
func deleteAssociation(p preparer, t table) (*sql.Stmt, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1 AND %s = $2;",
		t.TableName(),
		t.Columns()[0],
		t.Columns()[1],
	)

	return p.Prepare(query)
}

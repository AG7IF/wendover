package postgresrepo

import (
	"database/sql"
	"strings"

	"github.com/pkg/errors"

	"github.com/ag7if/wendover/internal/repositories"
)

func processDuplicateKey(object, key string, err error) (bool, error) {
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return true, repositories.ErrDuplicateKey{Object: object, Key: key}
	}
	return false, err
}

func processNotFound(object, key string, err error) (bool, error) {
	if errors.Is(err, sql.ErrNoRows) {
		return true, repositories.ErrNotFound{Object: object, Key: strings.ToUpper(key)}
	}
	return false, err
}

func processError(object, key string, err error) error {
	var match bool
	var e error

	if match, e = processNotFound(object, key, err); match {
		return e
	} else if match, e = processDuplicateKey(object, key, err); match {
		return e
	} else {
		return err
	}
}

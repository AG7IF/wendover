package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/internal/config"
)

type ErrDirtySchema struct {
	Version uint
}

func (eds ErrDirtySchema) Error() string {
	return fmt.Sprintf("database at version %d reports a dirty schema", eds.Version)
}

type Migration struct {
	migration *migrate.Migrate
}

func NewMigration() (*Migration, error) {
	migrationsPath := viper.GetString(config.DatabaseMigrationSource)
	dbURL, err := GetDBUrl()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Migration{migration: m}, nil
}

func (m *Migration) Up() (uint, error) {
	err := m.migration.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			version, _, _ := m.migration.Version()
			return version, nil
		}

		return 0, errors.WithStack(err)
	}

	version, dirty, err := m.migration.Version()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	if dirty {
		return 0, errors.WithStack(ErrDirtySchema{Version: version})
	}

	return version, nil
}

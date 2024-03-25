package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/internal/config"
)

func GetDBUrl() string {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		viper.GetString(config.DatabaseUser),
		viper.GetString(config.DatabasePassword),
		viper.GetString(config.DatabaseHost),
		viper.GetString(config.DatabasePort),
		viper.GetString(config.DatabaseName),
	)

	if !viper.GetBool(config.DatabaseSSL) {
		return fmt.Sprintf("%s?sslmode=disable", url)
	}

	return url
}

func GetDB() (*sql.DB, error) {
	url := GetDBUrl()
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

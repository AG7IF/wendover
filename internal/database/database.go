package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/internal/config"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetDBUrl() (string, error) {
	credentials := Credentials{}

	credStr := viper.GetString(config.DatabaseCredentials)
	err := json.Unmarshal([]byte(credStr), &credentials)
	if err != nil {
		return "", errors.WithStack(err)
	}

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		url.QueryEscape(credentials.Username),
		url.QueryEscape(credentials.Password),
		viper.GetString(config.DatabaseHost),
		viper.GetString(config.DatabasePort),
		viper.GetString(config.DatabaseName),
	)

	if !viper.GetBool(config.DatabaseSSL) {
		return fmt.Sprintf("%s?sslmode=disable", dbUrl), nil
	}

	return dbUrl, nil
}

func GetDB() (*sql.DB, error) {
	url, err := GetDBUrl()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

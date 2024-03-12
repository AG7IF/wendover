package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/romanyx/polluter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	_ "github.com/ag7if/wendover/internal/config"
	"github.com/ag7if/wendover/internal/database"
	"github.com/ag7if/wendover/internal/logging"
)

func reset(pollute bool) error {
	err := down()
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		return errors.WithStack(err)
	}
	err = up(pollute)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func down() error {
	log.Info().Str("database", viper.GetString("database.name")).Msg("migrating database down")
	m, _, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Down()
	if err != nil {
		return errors.WithStack(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return errors.WithStack(err)
	}
	log.Info().Str("database", viper.GetString("database.name")).Uint("version", version).Bool("dirty", dirty).Msg("migration complete")

	return nil
}

func up(pollute bool) error {
	log.Info().Str("database", viper.GetString("database.name")).Msg("migrating database up")
	m, p, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Up()
	if err != nil {
		return errors.WithStack(err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return errors.WithStack(err)
	}
	log.Info().Str("database", viper.GetString("database.name")).Uint("version", version).Bool("dirty", dirty).Msg("migration complete")

	if pollute {
		err = seed(p)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func seed(polluter *polluter.Polluter) error {
	if polluter == nil {
		_, p, err := setup()
		if err != nil {
			return errors.WithStack(err)
		}
		polluter = p
	}

	log.Info().Str("database", viper.GetString("database.name")).Msg("seeding database")
	seedPath := viper.GetString("database.migration.seed")

	f, err := os.Open(seedPath)
	if err != nil {
		log.Error().Err(err).Str("database", viper.GetString("database.name")).Msg("unable to read database seed file")
	}
	defer f.Close()

	err = polluter.Pollute(f)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func setup() (*migrate.Migrate, *polluter.Polluter, error) {
	dbURL := database.GetDBUrl()
	log.Debug().Str("url", dbURL).Msg("database connection created")

	migrationsPath := viper.GetString("database.migration.source")
	log.Debug().Str("path", migrationsPath).Msg("migrations path identified")

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	p := polluter.New(polluter.PostgresEngine(db), polluter.YAMLParser)

	return m, p, nil
}

func main() {
	loglevel := flag.String("loglevel", "info", "set the loglevel for the tool")
	pollute := flag.Bool("seed", false, "seed database with test data after migration")
	flag.Parse()

	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Println("Usage: go run migrate.go [--seed] [--loglevel=debug|info|warn|error] [up|down|reset|seed]")
		os.Exit(1)
	}

	logging.InitLogging(*loglevel, true)

	var err error
	switch strings.ToLower(cmd) {
	case "up":
		err = up(*pollute)
	case "down":
		err = down()
	case "reset":
		err = reset(*pollute)
	case "seed":
		err = seed(nil)
	default:
		fmt.Printf("Invalid command: %s\nValid commands are [up|down|reset]\n", cmd)
	}

	if err != nil {
		log.Error().Err(err).Str("command", cmd).Msg("problem encountered while running command")
		os.Exit(1)
	}

	os.Exit(0)
}

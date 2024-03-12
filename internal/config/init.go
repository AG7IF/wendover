package config

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	// Defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.name", "wendover_dev")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.ssl", false)
	viper.SetDefault("database.migration.source", "file:///path/to/migrations")
	viper.SetDefault("database.migration.seed", "/path/to/seed")

	// Config Directory
	cfgDir, err := CfgDir()
	if err != nil {
		log.Error().Err(err).Msg("failed to find config dir while initializing config")
	} else {
		log.Debug().Str("cfgDir", cfgDir).Msg("configuration directory found")
	}

	// Read in Config File
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(cfgDir, "cfg"))
	err = viper.ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "cfg", "config.yaml")
		log.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = viper.WriteConfigAs(path)
		if err != nil {
			log.Error().Err(err).Msg("failed to create default config file")
		}
	}

	// TODO: env overrides
}

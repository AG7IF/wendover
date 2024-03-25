package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	appName   = "wendover"
	envPrefix = "WENDOVER"
)

const (
	Version                 = "version"
	Directory               = "config.directory"
	DatabaseHost            = "database.host"
	DatabasePort            = "database.port"
	DatabaseName            = "database.name"
	DatabaseUser            = "database.user"
	DatabasePassword        = "database.password"
	DatabaseSSL             = "database.ssl"
	DatabaseMigrationSource = "database.migration.source"
	DatabaseMigrationSeed   = "database.migration.seed"
)

func initDefaults() {
	// Use UserConfigDir as the default for config file path
	usrCfgDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	// Defaults
	viper.SetDefault(Version, "")
	viper.SetDefault(Directory, filepath.Join(usrCfgDir, appName))
	viper.SetDefault(DatabaseHost, "localhost")
	viper.SetDefault(DatabasePort, "5432")
	viper.SetDefault(DatabaseName, "wendover_dev")
	viper.SetDefault(DatabaseUser, "postgres")
	viper.SetDefault(DatabasePassword, "postgres")
	viper.SetDefault(DatabaseSSL, false)
	viper.SetDefault(DatabaseMigrationSource, "file:///path/to/migrations")
	viper.SetDefault(DatabaseMigrationSeed, "/path/to/seed")
}

func initEnv() {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func initCfgFile() {
	cfgDir := viper.GetString(Directory)

	// Read in Config File
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(cfgDir, "cfg"))
	err := viper.ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "cfg", "config.yaml")
		log.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = viper.WriteConfigAs(path)
		if err != nil {
			log.Error().Err(err).Msg("failed to create default config file")
		}
	}
}

func init() {
	initDefaults()
	initEnv()

	// To skip initiation of the config file, set the WENDOVER_CONFIF_DIRECTORY environment variable to DOZFAC
	// (DOZen FACtors, i.e. set this up as a 12-factor application).
	if viper.GetString(Directory) != "DOZFAC" {
		initCfgFile()
	}
}

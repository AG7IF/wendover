package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/internal/aws"
)

const (
	AppName   = "wendover"
	envPrefix = "WENDOVER"
)

const (
	Version                 = "version"
	AWSRegion               = "aws.region"
	AWSCognitoIss           = "aws.cognito.iss"
	AWSCognitoUserpoolID    = "aws.cognito.userpool_id"
	AWSLogGroupName         = "aws.log_group_name"
	AWSLogStreamName        = "aws.log_stream_name"
	Directory               = "config.directory"
	DatabaseHost            = "database.host"
	DatabasePort            = "database.port"
	DatabaseName            = "database.name"
	DatabaseCredentials     = "database.credentials"
	DatabaseSSL             = "database.ssl"
	DatabaseMigrationSource = "database.migration.source"
	DatabaseMigrationSeed   = "database.migration.seed"
	ServerRunAddress        = "server.run_address"
	ServerRootPath          = "server.root_path"
)

func initDefaults() {
	// Use UserConfigDir as the default for config file path
	usrCfgDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	// Defaults
	viper.SetDefault(Version, "")
	viper.SetDefault(Directory, filepath.Join(usrCfgDir, AppName))
	viper.SetDefault(AWSRegion, "")
	viper.SetDefault(AWSCognitoIss, "")
	viper.SetDefault(AWSCognitoUserpoolID, "")
	viper.SetDefault(AWSLogGroupName, "/aws/lambda/wendover-migrate-db")
	viper.SetDefault(AWSLogStreamName, "wendover-migrate-db")
	viper.SetDefault(DatabaseHost, "localhost")
	viper.SetDefault(DatabasePort, "5432")
	viper.SetDefault(DatabaseName, "wendover_dev")
	viper.SetDefault(DatabaseCredentials, `{"username": "postgres", "password": "postgres"}`)
	viper.SetDefault(DatabaseSSL, false)
	viper.SetDefault(DatabaseMigrationSource, "file:///path/to/migrations")
	viper.SetDefault(DatabaseMigrationSeed, "/path/to/seed")
	viper.SetDefault(ServerRunAddress, "0.0.0.0:8080")
	viper.SetDefault(ServerRootPath, "/api/v1")
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
			log.Error().Stack().Err(err).Msg("failed to create default config file")
		}
	}
}

func init() {
	initDefaults()
	initEnv()

	switch viper.GetString(Directory) {
	case "DOZFAC":
		break
	case "AWS":
		aws.InitConfigFromParameterStore()
	default:
		initCfgFile()
	}
}

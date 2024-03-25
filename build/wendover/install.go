package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/ag7if/wendover/build"
	"github.com/ag7if/wendover/internal/config"
	"github.com/ag7if/wendover/internal/logging"
)

func CreateConfigDirectories() error {
	subdirs := []string{
		"assets",
		"cfg",
		"defs",
		"schemas",
	}

	cfgDir := viper.GetString(config.Directory)
	err := build.CreateDir(cfgDir)
	if err != nil {
		return errors.WithStack(err)
	}

	err = build.CreateSubDirs(cfgDir, subdirs)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func CreateCacheDirectories() error {
	panic("function not implemented")
}

func CopyAssets() error {
	panic("function not implemented")
}

func main() {
	logging.InitLogging("info", true)

	// projectDir := os.Args[1]

	err := CreateConfigDirectories()
	if err != nil {
		log.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	/*
		err = CreateCacheDirectories()
		if err != nil {
			log.Error().Err(err).Msg("failed to create cache directories")
			os.Exit(1)
		}

		err = CopyAssets()
		if err != nil {
			log.Error().Err(err).Msg("failed to copy assets")
			os.Exit(1)
		}
	*/
}

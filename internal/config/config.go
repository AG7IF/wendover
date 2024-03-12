package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	appName       = "wendover"
	baseConfigKey = "WENDOVER"
)

func CfgDir() (string, error) {
	cfgDir := os.Getenv(fmt.Sprintf("%s_CONFIG", baseConfigKey))
	if cfgDir != "" {
		return cfgDir, nil
	}

	hd, err := os.UserConfigDir()
	if err != nil {
		return "", errors.WithMessage(err, "failed to find user config directory")
	}

	return filepath.Join(hd, appName), nil
}

func CacheDir() (string, error) {
	cacheDir := os.Getenv(fmt.Sprintf("%s_CACHE", baseConfigKey))
	if cacheDir != "" {
		return cacheDir, nil
	}

	cd, err := os.UserCacheDir()
	if err != nil {
		return "", errors.WithMessage(err, "failed to find user cache directory")
	}

	return filepath.Join(cd, appName), nil
}

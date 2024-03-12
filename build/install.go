package build

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func clearFileExistsError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "file exists") {
		return nil
	}

	return err
}

func CreateDir(dir string) error {
	log.Info().Str("dir", dir).Msg("creating directory")
	err := os.Mkdir(dir, 0700)
	err = clearFileExistsError(err)
	if err != nil {
		return errors.WithMessagef(err, "failed to create dir: %s", dir)
	}

	return nil
}

func CreateSubDirs(dir string, subdirs []string) error {
	for _, v := range subdirs {
		path := filepath.Join(dir, v)
		err := CreateDir(path)
		if err != nil {
			return errors.WithMessagef(err, "unexpected error encountered while creating subdir: %s", path)
		}
	}

	return nil
}

func copyAsset() error {
	panic("function not implemented")
}

func CopyAssets(projectPath, destinationPath string) {
	panic("function not implemented")
}

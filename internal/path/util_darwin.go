//go:build darwin

package path

import (
	"os"
	"path/filepath"
)

func loginUsersVDF(steamDir string) (string, error) {
	if steamDir == "" {
		applicationSupport, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		steamDir = filepath.Join(applicationSupport, "Steam")
	}
	return filepath.Join(steamDir, "config", "loginusers.vdf"), nil
}

func localVDF(steamDir string) (string, error) {
	if steamDir == "" {
		applicationSupport, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		steamDir = filepath.Join(applicationSupport, "Steam")
	}
	return filepath.Join(steamDir, "local.vdf"), nil
}

//go:build linux

package store

import (
	"os"
	"path/filepath"
)

func loginUsersVDF(steamDir string) (string, error) {
	if steamDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		steamDir = filepath.Join(home, ".local", "share", "Steam")
	}
	return filepath.Join(steamDir, "config", "loginusers.vdf"), nil
}

func localVDF(steamDir string) (string, error) {
	if steamDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		steamDir = filepath.Join(home, ".local", "share", "Steam")
	}
	return filepath.Join(steamDir, "local.vdf"), nil
}

//go:build windows

package store

import (
	"os"
	"path/filepath"
)

func loginUsersVDF(steamDir string) (string, error) {
	if steamDir == "" {
		steamDir = filepath.Join("Program Files (x86)", "Steam")
	}
	return filepath.Join(steamDir, "config", "loginusers.vdf"), nil
}

func localVDF(_ string) (string, error) {
	localAppData, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(localAppData, "Steam", "local.vdf"), nil
}

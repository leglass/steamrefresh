//go:build !windows

package path

import "errors"

// Unsupported
func GetLoginUsers(steamDir string) (map[uint64][2]string, error) {
	return nil, errors.ErrUnsupported
}

// Unsupported
func GetConnectCache(steamDir string) (map[string]string, error) {
	return nil, errors.ErrUnsupported
}

//go:build !windows && !darwin && !linux

package path

import "errors"

// Unsupported
func localVDF(_ string) (string, error) {
	return "", errors.ErrUnsupported
}

// Unsupported
func loginUsersVDF(_ string) (string, error) {
	return "", errors.ErrUnsupported
}

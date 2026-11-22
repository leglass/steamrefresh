//go:build !windows

package crypto

import "errors"

// Unsupported
func Decrypt(entropy []byte, dataIn []byte) ([]byte, error) {
	return nil, errors.ErrUnsupported
}

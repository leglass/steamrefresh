//go:build !windows

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"fmt"
)

// Decryption via Valve CCrypto::SymmetricDecrypt scheme, with SHA-256 hashed accountName as AES-256 key
// The first 16 bytes of dataIn are the IV encrypted with AES-ECB, the remainder is the token encrypted with AES-CBC.
func Decrypt(accountName []byte, dataIn []byte) ([]byte, error) {
	key := sha256.Sum256(accountName)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("steamrefresh: new cipher: %w", err)
	}

	if len(dataIn) < 2*aes.BlockSize || len(dataIn)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("steamrefresh: bad ciphertext length %d", len(dataIn))
	}

	iv := make([]byte, aes.BlockSize)
	block.Decrypt(iv, dataIn[:aes.BlockSize])

	body := dataIn[aes.BlockSize:]
	out := make([]byte, len(body))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(out, body)

	return pkcs7Unpad(out, aes.BlockSize)
}

func pkcs7Unpad(b []byte, blockSize int) ([]byte, error) {
	n := len(b)
	if n == 0 || n%blockSize != 0 {
		return nil, errors.New("steamrefresh: invalid padded length")
	}
	pad := int(b[n-1])
	if pad == 0 || pad > blockSize {
		return nil, errors.New("steamrefresh: invalid padding")
	}
	for _, c := range b[n-pad:] {
		if int(c) != pad {
			return nil, errors.New("steamrefresh: invalid padding")
		}
	}
	return b[:n-pad], nil
}

package steamrefresh

import (
	"encoding/hex"
	"errors"
	"fmt"
	"hash/crc32"
	"io/fs"

	"github.com/leglass/steamrefresh/internal/crypto"
	"github.com/leglass/steamrefresh/internal/path"
)

type Client struct {
	// Defaults to:
	// Windows:  /Program Files (x86)/Steam
	// Darwin:   ~/Library/Application Support/Steam
	// Linux:    ~/.local/share/Steam
	SteamDir string

	// Defaults to:
	// Windows:  ~/AppData/Local/Steam
	// MacOS:    ~/Library/Application Support/Steam
	// Linux:    ~/.local/share/Steam
	SteamCacheDir string
}

type User struct {
	AccountName string
	PersonaName string // Profile name
	Steam64     uint64 // Steam64 ID
}

// Returns slice of cached User(s) in loginusers.vdf
// Typically only one user is present at a time
func (c *Client) UserList() ([]User, error) {
	entries, err := path.GetLoginUsers(c.SteamDir)
	if err != nil {
		if _, ok := errors.AsType[*fs.PathError](err); ok {
			return nil, ErrNoLoginUsers
		}
		return nil, err
	}

	i := 0
	users := make([]User, len(entries))
	for k, v := range entries {
		users[i] = User{
			AccountName: v[0],
			PersonaName: v[1],
			Steam64:     k,
		}
		i++
	}

	return users, nil
}

// Returns cached refresh token for accountName if present, returns empty string if not present
func (c *Client) Token(accountName string) (string, error) {
	cc, err := path.GetConnectCache(c.SteamCacheDir)
	if err != nil {
		if _, ok := errors.AsType[*fs.PathError](err); ok {
			return "", ErrNoLocalCache
		}
		return "", err
	}

	checksum := fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(accountName)))

	for u, v := range cc {
		if u[0:8] == checksum {
			dataIn, err := hex.DecodeString(v)
			if err != nil {
				return "", err
			}
			dataOut, err := crypto.Decrypt([]byte(accountName), []byte(dataIn))
			if err != nil {
				return "", err
			}
			return string(dataOut), nil
		}
	}

	return "", nil
}

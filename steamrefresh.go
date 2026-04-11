package steamrefresh

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"

	"github.com/leglass/steamrefresh/crypto"
	"github.com/leglass/steamrefresh/path"
	"github.com/leglass/steamrefresh/types"
)

type Client struct {
	SteamDir string // Defaults to %ProgramFiles(X86)%/Steam/ on Windows if empty
}

// Returns slice of cached Users in loginusers.vdf
func (c *Client) GetUsers() ([]types.User, error) {
	users, err := path.GetLoginUsers(c.SteamDir)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Takes in user AccountName, returns AutoLogin Refresh Token if available, returns empty string if not available
func (c *Client) GetTokenForUser(user string) (string, error) {
	cc, err := path.GetConnectCache(c.SteamDir)
	if err != nil {
		return "", err
	}

	checksum := fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(user)))

	for u, v := range cc {
		if u[0:8] == checksum {
			dataIn, err := hex.DecodeString(v.(string))
			if err != nil {
				return "", err
			}
			dataOut, err := crypto.Decrypt([]byte(user), []byte(dataIn))
			if err != nil {
				return "", err
			}
			return string(dataOut), nil
		}
	}

	return "", nil
}

# steamrefresh

A lightweight Go package for reading the locally cached auto-login *refresh tokens* from an installed copy of the Steam Client.

It parses Steam's registry of active account(s) in `loginusers.vdf`, then decrypts the associated cached JWT in `local.vdf`.

## Install

```sh
go get github.com/leglass/steamrefresh
```

## Usage

```go
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/leglass/steamrefresh"
)

func main() {
	// SteamDir empty = use the platform default (see below).
	c := &steamrefresh.Client{}

	users, err := c.UserList()
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		token, err := c.Token(u.AccountName)
		if errors.Is(err, steamrefresh.ErrNoTokenEntry) {
			fmt.Printf("%s (%d): no cached token\n", u.AccountName, u.Steam64)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s (%d): %s\n", u.AccountName, u.Steam64, token)
	}
}
```

## API

```go
type Client struct {
	SteamDir string // Steam install directory; empty = platform default
}

type User struct {
	AccountName string
	PersonaName string // profile display name
	Steam64     uint64 // SteamID64
}

func (c *Client) UserList() ([]User, error)        // cached accounts from loginusers.vdf
func (c *Client) Token(accountName string) (string, error) // refresh token for an account
```

### Errors

| Error | Description |
|---|---|
| `ErrNoLoginUsers` | `loginusers.vdf` could not be read at `SteamDir` |
| `ErrNoLocalCache` | `local.vdf` could not be read at the cache directory |
| `ErrNoTokenEntry` | no cached token entry exists for the given `accountName` |

## Platform support

| OS | Scheme | Default SteamDir |
|---|---|---|
| Windows | DPAPI (`CryptUnprotectData`) | `C:\Program Files (x86)\Steam` |
| macOS | AES-256 (Valve `CCrypto::SymmetricDecrypt`) | `~/Library/Application Support/Steam` |
| Linux | AES-256 (Valve `CCrypto::SymmetricDecrypt`) | `~/.local/share/Steam` |

On Windows, `local.vdf` is read from the user cache directory (`%LocalAppData%\Steam`) instead of the Steam install directory.

## Notes

- Since DPAPI uses a per-user credential store, only JWTs cached by the current Windows user can be decrypted. On all other platforms decryption is user-independent.
- In modern versions of the steam client, only one account is present in loginusers.vdf/local.vdf at a time.

## License

[MIT](LICENSE)

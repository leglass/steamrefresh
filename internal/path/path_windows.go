//go:build windows

package path

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/andygrunwald/vdf"
)

// Returns {AccountName, PersonaName} mapped to SteamID64
func GetLoginUsers(steamDir string) (map[uint64][2]string, error) {
	users := make(map[uint64][2]string)

	if steamDir == "" {
		steamDir = filepath.Join(os.Getenv("ProgramFiles(X86)"), "Steam")
	}

	file, err := os.Open(filepath.Join(steamDir, "config", "loginusers.vdf"))
	if err != nil {
		return nil, err
	}

	p := vdf.NewParser(file)
	m, err := p.Parse()
	if err != nil {
		return nil, err
	}

	for k, v := range m["users"].(map[string]any) {
		steamid, err := strconv.ParseUint(k, 10, 64)
		if err != nil {
			return nil, err
		}
		users[steamid] = [2]string{v.(map[string]any)["AccountName"].(string), v.(map[string]any)["PersonaName"].(string)}
	}

	return users, nil
}

// Obtains ConnectCache entries from local.vdf
func GetConnectCache(cacheDir string) (map[string]string, error) {
	if cacheDir == "" {
		cacheDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "Steam")
	}

	file, err := os.Open(filepath.Join(cacheDir, "local.vdf"))
	if err != nil {
		return nil, err
	}

	p := vdf.NewParser(file)
	m, err := p.Parse()
	if err != nil {
		return nil, err
	}

	return toStringMap(m["MachineUserConfigStore"].(map[string]any)["Software"].(map[string]any)["Valve"].(map[string]any)["Steam"].(map[string]any)["ConnectCache"].(map[string]any)), nil
}

package steamrefresh

import "errors"

var (
	// Returned when no entry exists for a given accountName in local.vdf
	ErrNoTokenEntry = errors.New("steamrefresh: no cached token available")

	// ErrNoLoginUsers is returned when loginusers.vdf is unable to be read.
	ErrNoLoginUsers = errors.New("steamrefresh: loginusers.vdf is unable to be read at given SteamDir")

	// ErrNoLoginUsers is returned when local.vdf is unable to be read.
	ErrNoLocalCache = errors.New("steamrefresh: local.vdf is unable to be read at cache directory")
)

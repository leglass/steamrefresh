package types

type User struct {
	AccountName string
	PersonaName string // Profile name
	Steam64     string // Steam64 ID
	AutoLogin   bool   // Whether user is likely to have a valid Connect Cache JWT saved
}

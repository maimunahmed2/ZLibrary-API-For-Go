package types

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenLoginIn struct {
	RemixUserId  string `json:"remixUserId"`
	RemixUserKey string `json:"remixUserKey"`
}
package token

import "time"

// Maker is an interface for managing jwt tokens
type Maker interface {
	// CreateToken creates a new token for a specific id and userName and expireDuration
	CreateToken(id int64, userName string, expireDuration time.Duration) (string, error)
	// VerifyToken checks if the token is valid or not
	VerifyToken(myToken string) (string, int64, bool)
	// ParseUserIdByToken ParseUserIdByToken
	ParseUserIdByToken(tokenString string) (userId int64, err error)
}

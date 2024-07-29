package token

import "time"

// Maker is an interface for managin tokens
type Maker interface {

	// CreateToken creates a new token for an specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	//  CerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

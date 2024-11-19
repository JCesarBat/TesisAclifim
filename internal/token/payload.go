package token

import (
	"TesisAclifim/pkg/util"
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// Different types of error returned by the VerifyToken
var (
	ErrInvalidToken = errors.New("Invalid Token")
	ErrExpiredToken = errors.New("Token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func randomInt64(min int64, max int64) int64 {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	delta := max - min + 1
	value := generator.Int63n(int64(delta))

	return value + min
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID := util.RandomUUID()

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

package util

import "github.com/google/uuid"

func RandomUUID() uuid.UUID {
	randomUUID := uuid.New()
	return randomUUID
}

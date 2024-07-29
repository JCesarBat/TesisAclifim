// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Municipio struct {
	ID          int64  `json:"id"`
	IDProvincia int64  `json:"id_provincia"`
	Name        string `json:"name"`
}

type Provincium struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID          int64        `json:"id"`
	IDMunicipio int64        `json:"id_municipio"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	SuperUser   sql.NullBool `json:"super_user"`
	CreatedAt   time.Time    `json:"created_at"`
}

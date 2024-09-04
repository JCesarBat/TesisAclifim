package Prov

import (
	database "Tesis/database/sqlc"
)

type ProvInterface interface {
	GetStore() database.Store
}

type Server struct {
	ProvInterface
}

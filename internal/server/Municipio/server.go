package Municipio

import (
	database "Tesis/database/sqlc"
)

type MunInterface interface {
	GetStore() database.Store
}

type Server struct {
	MunInterface
}

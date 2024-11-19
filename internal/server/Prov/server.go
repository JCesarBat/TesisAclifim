package Prov

import database "TesisAclifim/database/sqlc"

type ProvInterface interface {
	GetStore() database.Store
}

type Server struct {
	ProvInterface
}

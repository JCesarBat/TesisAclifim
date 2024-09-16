package orchestrator

import (
	database "Tesis/database/sqlc"
	"Tesis/internal/server/Municipio"
	"Tesis/internal/server/Prov"
	"Tesis/internal/server/asociado"
	"Tesis/internal/server/auth"
	"Tesis/internal/server/common_data"
	"Tesis/internal/server/datos_sociales"
	"Tesis/internal/server/users"
	"Tesis/pkg/util"
)

type Orchestrator struct {
	Mun      *Municipio.Server
	Prov     *Prov.Server
	Auth     *auth.Server
	User     *users.Server
	Asociado *asociado.Server
	DatosS   *datos_sociales.Server
}

func NewOrchestrator(store database.Store, config util.Config) (*Orchestrator, error) {
	server, err := common_data.NewServer(store, config)
	if err != nil {
		return nil, err
	}
	return &Orchestrator{
		Mun:      &Municipio.Server{server},
		Prov:     &Prov.Server{server},
		Auth:     &auth.Server{server},
		User:     &users.Server{server},
		Asociado: &asociado.Server{server},
		DatosS:   &datos_sociales.Server{server},
	}, nil
}

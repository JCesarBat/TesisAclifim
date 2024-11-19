package orchestrator

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/Municipio"
	"TesisAclifim/internal/server/Prov"
	"TesisAclifim/internal/server/actividad_cultural"
	"TesisAclifim/internal/server/actividad_deportiva"
	"TesisAclifim/internal/server/actividad_educativa"
	"TesisAclifim/internal/server/asociado"
	"TesisAclifim/internal/server/auth"
	"TesisAclifim/internal/server/common_data"
	"TesisAclifim/internal/server/datos_sociales"
	"TesisAclifim/internal/server/users"
	"TesisAclifim/pkg/util"
)

type Orchestrator struct {
	Mun      *Municipio.Server
	Prov     *Prov.Server
	Auth     *auth.Server
	User     *users.Server
	Asociado *asociado.Server
	DatosS   *datos_sociales.Server
	ActDepor *actividad_deportiva.Server
	ActEduc  *actividad_educativa.Server
	ActCult  *actividad_cultural.Server
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
		ActEduc:  &actividad_educativa.Server{server},
		ActDepor: &actividad_deportiva.Server{server},
		ActCult:  &actividad_cultural.Server{server},
	}, nil
}

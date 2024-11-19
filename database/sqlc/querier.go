// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CreateActividadCultural(ctx context.Context, arg CreateActividadCulturalParams) (ActividadCultural, error)
	CreateActividadDeportiva(ctx context.Context, arg CreateActividadDeportivaParams) (ActividadDeportiva, error)
	CreateActividadEducativa(ctx context.Context, arg CreateActividadEducativaParams) (ActividadEducativa, error)
	CreateEstudiosActuales(ctx context.Context, arg CreateEstudiosActualesParams) (EstudiosActuale, error)
	CreateParticipacionCultural(ctx context.Context, arg CreateParticipacionCulturalParams) (ParticipacionC, error)
	CreateParticipacionD(ctx context.Context, arg CreateParticipacionDParams) (ParticipacionD, error)
	CreateRol(ctx context.Context, rol string) (Rol, error)
	CreateUSessions(ctx context.Context, arg CreateUSessionsParams) (Session, error)
	DeleteActividaDeportiva(ctx context.Context, id int64) error
	DeleteActividadCulural(ctx context.Context, id int64) error
	DeleteActividadEducativa(ctx context.Context, id int64) error
	DeleteAsociado(ctx context.Context, id int64) error
	DeleteDatosSociales(ctx context.Context, idAsociado int64) error
	DeleteEstudiosActuales(ctx context.Context, id int64) error
	DeleteParticipacion(ctx context.Context, id int64) error
	DeleteParticipacionC(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetActividadCultural(ctx context.Context, id int64) (ActividadCultural, error)
	GetActividadDeportiva(ctx context.Context, id int64) (ActividadDeportiva, error)
	GetActividadEducativa(ctx context.Context, id int64) (ActividadEducativa, error)
	GetAllMunicipio(ctx context.Context, idProvincia int64) ([]Municipio, error)
	GetAllProv(ctx context.Context) ([]Provincium, error)
	GetAllRol(ctx context.Context) ([]Rol, error)
	GetAllRolFromUser(ctx context.Context, userID int64) ([]Rol, error)
	GetAsociado(ctx context.Context, id int64) (Asociado, error)
	GetDatosSociales(ctx context.Context, idAsociado int64) (DatosSociale, error)
	GetEstudiosActuales(ctx context.Context, id int64) (EstudiosActuale, error)
	GetMunicipio(ctx context.Context, id int64) (Municipio, error)
	GetParticipacionC(ctx context.Context, idActividadCultural sql.NullInt64) ([]ParticipacionC, error)
	GetParticipacionD(ctx context.Context, idActividadDeportiva sql.NullInt64) ([]ParticipacionD, error)
	GetProvincia(ctx context.Context, id int64) (Provincium, error)
	GetSessions(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserID(ctx context.Context, id int64) (User, error)
	InertarProv(ctx context.Context, name string) (Provincium, error)
	InsertAsoiciado(ctx context.Context, arg InsertAsoiciadoParams) (Asociado, error)
	InsertDatosSociales(ctx context.Context, arg InsertDatosSocialesParams) (DatosSociale, error)
	InsertMunicipio(ctx context.Context, arg InsertMunicipioParams) (Municipio, error)
	InsertRolUser(ctx context.Context, arg InsertRolUserParams) (UserRole, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	ListAsociado(ctx context.Context, arg ListAsociadoParams) ([]Asociado, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateActividadCultural(ctx context.Context, arg UpdateActividadCulturalParams) (ActividadCultural, error)
	UpdateActividadDeportiva(ctx context.Context, arg UpdateActividadDeportivaParams) (ActividadDeportiva, error)
	UpdateActividadEducativa(ctx context.Context, arg UpdateActividadEducativaParams) (ActividadEducativa, error)
	UpdateAsociado(ctx context.Context, arg UpdateAsociadoParams) (Asociado, error)
	UpdateDatosSociales(ctx context.Context, arg UpdateDatosSocialesParams) (DatosSociale, error)
	UpdateEstudiosActuales(ctx context.Context, arg UpdateEstudiosActualesParams) (EstudiosActuale, error)
	UpdateParticipacionC(ctx context.Context, arg UpdateParticipacionCParams) (ParticipacionC, error)
	UpdateParticipacionD(ctx context.Context, arg UpdateParticipacionDParams) (ParticipacionD, error)
	UpdateToSuperUser(ctx context.Context, arg UpdateToSuperUserParams) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)

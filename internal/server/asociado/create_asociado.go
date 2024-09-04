package asociado

import (
	database "Tesis/database/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type CreateAsociadoRequest struct {
	Name                string `json:"name" binding:"required"`
	Apellido1           string `json:"Apellido" binding:"required"`
	Apellido2           string `json:"Apellido"binding:"required"`
	Activo              bool   `json:"Activo" binding:"required"`
	Carnet              int64  `json:"Carnet" binding:"required"`
	Sexo                bool   `json:"Sexo" binding:"required"`
	NumeroT             int64  `json:"NumeroT"`
	NumeroPerteneciente string `json:"NumeroPerteneciente"`
	Direccion           string `json:"Direccion" binding:"required"`
	IDMunicipio         int64  `json:"IDMunicipio" binding:"required"`
}

func (server *Server) CreateAsociado(c *gin.Context) {
	var asociado CreateAsociadoRequest
	if err := c.ShouldBindJSON(&asociado); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	param := database.InsertAsoiciadoParams{
		Nombre:              asociado.Name,
		Apellido1:           asociado.Apellido1,
		Apellido2:           asociado.Apellido2,
		Activo:              asociado.Activo,
		Carnet:              asociado.Carnet,
		Sexo:                asociado.Sexo,
		NumeroT:             sql.NullInt64{Int64: asociado.NumeroT, Valid: true},
		NumeroPerteneciente: sql.NullString{String: asociado.NumeroPerteneciente, Valid: true},
		Direccion:           asociado.Direccion,
		IDMunicipio:         asociado.IDMunicipio,
	}
	result, err := server.GetStore().InsertAsoiciado(c, param)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

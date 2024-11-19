package asociado

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateAsociadoRequest struct {
	id_asociado         int64  `json:"id" binding:"required"`
	Name                string `json:"name" `
	Apellido1           string `json:"Apellido" `
	Apellido2           string `json:"Apellido"`
	Activo              bool   `json:"Activo" `
	Carnet              int64  `json:"Carnet" `
	Sexo                bool   `json:"Sexo" `
	NumeroT             int64  `json:"NumeroT"`
	NumeroPerteneciente string `json:"NumeroPerteneciente"`
	Direccion           string `json:"Direccion" `
	IDMunicipio         int64  `json:"IDMunicipio" `
}

func (server *Server) UpdateAsociado(c *gin.Context) {
	var req UpdateAsociadoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	params := database.UpdateAsociadoParams{
		ID:                  req.id_asociado,
		Nombre:              req.Name,
		Apellido1:           req.Apellido1,
		Apellido2:           req.Apellido2,
		Activo:              req.Activo,
		Carnet:              req.Carnet,
		Sexo:                req.Sexo,
		NumeroT:             sql.NullInt64{Int64: req.NumeroT, Valid: true},
		NumeroPerteneciente: sql.NullString{String: req.NumeroPerteneciente, Valid: true},
		Direccion:           req.Direccion,
		IDMunicipio:         req.IDMunicipio,
	}
	_, err := server.GetStore().UpdateAsociado(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Asociado updated successfully"})

}

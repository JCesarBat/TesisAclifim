package datos_sociales

import (
	database "Tesis/database/sqlc"
	"Tesis/internal/server/common_data"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InsertDatosSocialesRequest struct {
	ID_asociado                int64  `json:"id"biding:"required"`
	Ocupacion                  string `json:"occupacion"biding:"required"`
	EstadoCivil                string `json:"estadoCivil"biding:"required"`
	Integracion_Revolucionaria string `json:"integracionRevolucionaria"biding:"required"`
}

func (s *Server) InsertDatosSociales(c *gin.Context) {
	var req InsertDatosSocialesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	param := database.InsertDatosSocialesParams{
		IDAsociado:                req.ID_asociado,
		Ocupacion:                 sql.NullString{String: req.Ocupacion, Valid: true},
		EstadoCivil:               sql.NullString{String: req.EstadoCivil, Valid: true},
		IntegracionRevolucionaria: sql.NullString{String: req.Integracion_Revolucionaria, Valid: true},
	}
	datos, err := s.GetStore().InsertDatosSociales(c, param)
	if err != nil {
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}
	c.JSON(200, datos)
}

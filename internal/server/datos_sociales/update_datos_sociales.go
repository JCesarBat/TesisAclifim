package datos_sociales

import (
	database "Tesis/database/sqlc"
	"Tesis/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateDRequest struct {
	ID_asociado                int64  `json:"id"biding:"required"`
	Ocupacion                  string `json:"occupacion"`
	EstadoCivil                string `json:"estadoCivil"`
	Integracion_Revolucionaria string `json:"integracionRevolucionaria"`
}

func (s *Server) UpdateD(c *gin.Context) {
	var req UpdateDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	param := database.UpdateDatosSocialesParams{
		IDAsociado:                req.ID_asociado,
		Ocupacion:                 sql.NullString{String: req.Ocupacion, Valid: true},
		EstadoCivil:               sql.NullString{String: req.EstadoCivil, Valid: true},
		IntegracionRevolucionaria: sql.NullString{String: req.Integracion_Revolucionaria, Valid: true},
	}
	datos, err := s.GetStore().UpdateDatosSociales(c, param)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, datos)
}

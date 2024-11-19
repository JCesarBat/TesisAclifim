package actividad_deportiva

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"github.com/gin-gonic/gin"
)

type CreateActividadDeportivaRequest struct {
	Asociado_id int64    `json:"asociado_Id" biding:"required"`
	Aficcion    []string `json:"aficcion" biding:"required"`
}

func (s *Server) CreateAD(c *gin.Context) {
	var req CreateActividadDeportivaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, common_data.ErrorResponse(err))
		return
	}
	param := database.CreateActividadDeportivaParams{
		IDAsociado:        req.Asociado_id,
		AficcionOPractica: req.Aficcion,
	}
	actividad, err := s.GetStore().CreateActividadDeportiva(c, param)
	if err != nil {
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}
	c.JSON(200, actividad)
}

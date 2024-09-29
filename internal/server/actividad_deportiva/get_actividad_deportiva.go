package actividad_deportiva

import "C"
import (
	database "Tesis/database/sqlc"
	"Tesis/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetActividadDRequest struct {
	Asiociado_Id int64 `uri:"asiociado_id"biding:"required"`
}
type GetActividadDResponse struct {
	Asociado_id   int64                     `json:"asociado_Id" biding:"required"`
	Aficcion      []string                  `json:"aficcion" biding:"required"`
	Participacion []database.ParticipacionD `json:"participacion" biding:"required"`
}

func (s *Server) GetActividadDeportiva(c *gin.Context) {
	var req GetActividadDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, common_data.ErrorResponse(err))
		return
	}
	actividadD, err := s.GetStore().GetActividadDeportiva(c, req.Asiociado_Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}
	participacion, err := s.GetStore().GetParticipacionD(c, sql.NullInt64{Int64: actividadD.ID, Valid: true})
	if err != nil {
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}
	response := GetActividadDResponse{
		Asociado_id:   actividadD.IDAsociado,
		Aficcion:      actividadD.AficcionOPractica,
		Participacion: participacion,
	}
	c.JSON(200, response)
}

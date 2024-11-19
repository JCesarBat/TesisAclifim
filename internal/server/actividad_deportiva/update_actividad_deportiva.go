package actividad_deportiva

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateActividadDeportivaRequest struct {
	ID_asociado int64  `json:"ID_Asociado" biding:"required"`
	Aficcion    string `json:"aficcion" biding:"required"`
}

func (s *Server) UpdateActividadDeportiva(c *gin.Context) {
	var req UpdateActividadDeportivaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, common_data.ErrorResponse(err))
		return
	}
	param := database.UpdateActividadDeportivaParams{
		IDAsociado:  req.ID_asociado,
		ArrayAppend: req.Aficcion,
	}
	actividad, err := s.GetStore().UpdateActividadDeportiva(c, param)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}
	c.JSON(200, actividad)
}

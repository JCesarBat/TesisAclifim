package actividad_deportiva

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UpdateParticipacionDeportivaRequest struct {
	ID                    int64     `json:"ID" biding:"required"`
	ID_ActividadDeportiva int64     `json:"id_actividad_deportiva" biding:"required"`
	Deporte               string    `json:"deporte"`
	Fecha                 time.Time `json:"fecha"`
	LugarAlcanzado        int32     `json:"lugarAlcanzado"`
	DondeSeDesarrollo     string    `json:"dondeDesarrollo"`
}

func (s *Server) UpdatePorticipacionDeportiva(c *gin.Context) {
	var req UpdateParticipacionDeportivaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, common_data.ErrorResponse(err))
		return
	}
	param := database.UpdateParticipacionDParams{
		ID:                   req.ID,
		IDActividadDeportiva: sql.NullInt64{Int64: req.ID_ActividadDeportiva, Valid: true},
		Deporte:              req.Deporte,
		Fecha:                req.Fecha,
		LugarAlcanzado:       sql.NullInt32{Int32: req.LugarAlcanzado, Valid: true},
		DondeSeDesarrollo:    req.DondeSeDesarrollo,
	}
	participacion, err := s.GetStore().UpdateParticipacionD(c, param)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}

	c.JSON(200, participacion)

}

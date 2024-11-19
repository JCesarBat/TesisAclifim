package actividad_deportiva

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateParticipacionDRequest struct {
	Id_actividad_deportiva int64     `json:"id_actividad_deportiva"biding:"required"`
	Deporte                string    `json:"deporte"biding:"required"`
	fecha                  time.Time `json:"fecha"biding:"required"`
	Lugar_alcanzado        int32     `json:"lugar_alcanzado"biding:"required"`
	Donde_se_desarrollo    string    `json:"donde_se_desarrollo"biding:"required"`
}

func (s *Server) CreateParticipacionD(c *gin.Context) {
	var req CreateParticipacionDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, common_data.ErrorResponse(err))
		return
	}
	param := database.CreateParticipacionDParams{
		IDActividadDeportiva: sql.NullInt64{Int64: req.Id_actividad_deportiva, Valid: true},
		Deporte:              req.Deporte,
		Fecha:                req.fecha,
		LugarAlcanzado:       sql.NullInt32{Int32: req.Lugar_alcanzado, Valid: true},
		DondeSeDesarrollo:    req.Donde_se_desarrollo,
	}
	participacion, err := s.GetStore().CreateParticipacionD(c, param)
	if err != nil {
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}

	c.JSON(200, participacion)

}

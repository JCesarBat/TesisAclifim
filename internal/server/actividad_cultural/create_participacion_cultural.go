package actividad_cultural

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateParticipacionCulturalRequest struct {
	ID_actividad_cultural int64     `json:"id_actividad_cultural" binding:"required"`
	Especialidad          string    `json:"e_specialidad" binding:"required"`
	Fecha                 time.Time `json:"fecha" binding:"required"`
	LugarAlcanzado        int32     `json:"lugar_alcanzado" `
	Donde_se_desarrollo   string    `json:"donde_se_desarrollo" binding:"required"`
}

func (s *Server) CreateParticipacion(c *gin.Context) {
	var req CreateParticipacionCulturalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.CreateParticipacionCulturalParams{
		IDActividadCultural: sql.NullInt64{Int64: req.ID_actividad_cultural, Valid: true},
		Especialidad:        req.Especialidad,
		Fecha:               req.Fecha,
		LugarAlcanzado:      sql.NullInt32{Int32: req.LugarAlcanzado, Valid: true},
		DondeSeDesarrollo:   req.Donde_se_desarrollo,
	}
	part, err := s.GetStore().CreateParticipacionCultural(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, part)
}

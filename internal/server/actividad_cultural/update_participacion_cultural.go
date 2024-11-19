package actividad_cultural

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UpdateParticipacionCulturalRequest struct {
	ID                  int64     `json:"id" binding:"required"`
	Especialidad        string    `json:"e_specialidad" `
	Fecha               time.Time `json:"fecha" `
	LugarAlcanzado      int32     `json:"lugar_alcanzado" `
	Donde_se_desarrollo string    `json:"donde_se_desarrollo" `
}

func (s *Server) UpdateParticipacionCultural(c *gin.Context) {
	var req UpdateParticipacionCulturalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.UpdateParticipacionCParams{
		ID:                req.ID,
		Especialidad:      req.Especialidad,
		Fecha:             req.Fecha,
		LugarAlcanzado:    sql.NullInt32{Int32: req.LugarAlcanzado, Valid: true},
		DondeSeDesarrollo: req.Donde_se_desarrollo,
	}
	par, err := s.GetStore().UpdateParticipacionC(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, par)
}

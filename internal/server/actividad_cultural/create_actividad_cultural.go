package actividad_cultural

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateActividadCulturaRequest struct {
	Asociado_id  int64    `json:"asociado_Id" biding:"required"`
	Especialidad []string `json:"especialidad" biding:"required"`
}

func (s *Server) CreateActividadCultural(c *gin.Context) {
	var req CreateActividadCulturaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.CreateActividadCulturalParams{
		IDAsociado:   req.Asociado_id,
		Especialidad: req.Especialidad,
	}
	act, err := s.GetStore().CreateActividadCultural(c, param)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, act)
}

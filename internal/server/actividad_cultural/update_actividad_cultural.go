package actividad_cultural

import (
	database "TesisAclifim/database/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateActividadCulturalRequest struct {
	ID       int64  `json:"id" biding:"required"`
	Aficcion string `json:"aficcion" biding:"required"`
}

func (s *Server) UpdateActividadCultural(c *gin.Context) {
	var req UpdateActividadCulturalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.UpdateActividadCulturalParams{
		ID:          req.ID,
		ArrayAppend: req.Aficcion,
	}
	act, err := s.GetStore().UpdateActividadCultural(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, act)

}

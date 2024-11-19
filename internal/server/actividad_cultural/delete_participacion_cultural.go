package actividad_cultural

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteParticipacionCulturalR struct {
	ID int64 `json:"id" binding:"required"`
}

func (s *Server) DeleteParticipacionCultural(c *gin.Context) {
	var req DeleteParticipacionCulturalR
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.GetStore().DeleteParticipacionC(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": req.ID})
}

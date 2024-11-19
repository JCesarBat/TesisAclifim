package actividad_educativa

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteEstudiosActuales struct {
	ID int64 `json:"id" biding:"required"`
}

func (s *Server) DeleteEstudiosActuales(c *gin.Context) {
	var req DeleteEstudiosActuales
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.GetStore().DeleteEstudiosActuales(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

package asociado

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteAsociadoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) DeleteAsociado(c *gin.Context) {
	var request DeleteAsociadoRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete the asociado from the database
	if err := server.GetStore().DeleteAsociado(c, request.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asociado deleted"})
}

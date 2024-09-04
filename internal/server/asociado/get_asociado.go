package asociado

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAsociadoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) GetAsociado(c *gin.Context) {
	var req GetAsociadoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asociado, err := server.GetStore().GetAsociado(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asociado not found"})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, asociado)
}

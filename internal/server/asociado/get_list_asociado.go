package asociado

import (
	database "TesisAclifim/database/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListAsociadoRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListAsociados(c *gin.Context) {
	var req ListAsociadoRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.ListAsociadoParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	asociados, err := server.GetStore().ListAsociado(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": asociados})
}

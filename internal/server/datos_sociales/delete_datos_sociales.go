package datos_sociales

import (
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteDatosSociales struct {
	Asociado_id int64 `uri:"asociado_id" binding:"required"`
}

func (s *Server) DeleteDatosSociales(c *gin.Context) {
	var req DeleteDatosSociales
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	err := s.GetStore().DeleteDatosSociales(c, req.Asociado_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}

	c.JSON(200, gin.H{"message": "successfully"})
}

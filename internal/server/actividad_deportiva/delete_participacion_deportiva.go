package actividad_deportiva

import (
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DParticipacionD struct {
	ID int64 `uri:"id"biding:"required"`
}

func (s *Server) DeleteParticipacionDeportiva(c *gin.Context) {
	var req DParticipacionD
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}

	err := s.GetStore().DeleteParticipacion(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return

		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

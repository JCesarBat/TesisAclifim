package actividad_educativa

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetActividadEducativaRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetActividadEducativaResponse struct {
	database.ActividadEducativa `json:"actividad_educativa"`
	database.EstudiosActuale    `json:"actividad_educativa"`
}

func (s *Server) GetActividadEducativa(c *gin.Context) {
	var req GetActividadEducativaRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	act, err := s.GetStore().GetActividadEducativa(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	est, err := s.GetStore().GetEstudiosActuales(c, act.IDEstudiosActuales)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetActividadEducativaResponse{
		act,
		est,
	})
}

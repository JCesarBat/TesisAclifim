package actividad_educativa

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateActividadEducativaRequest struct {
	ID                    int64  `json:"id" binding:"required"`
	Ultimo_grado_aprobado string `json:"ultimo_grado_aprobado" binding:"required"`
}

func (s *Server) UpdateActividadEducativa(c *gin.Context) {
	var request UpdateActividadEducativaRequest
	var UltimoG database.UltimoGrado
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := UltimoG.Scan(request.Ultimo_grado_aprobado)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.UpdateActividadEducativaParams{
		ID:                  request.ID,
		UltimoGradoAprobado: UltimoG,
	}

	act, err := s.GetStore().UpdateActividadEducativa(c, param)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "no actividad educativa found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, act)
}

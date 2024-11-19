package actividad_educativa

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UpdateEstudiosActuales struct {
	Tipo_enseñansa       string    `json:"tipo_ense" `
	Centro               string    `json:"centro" `
	EspecialidadGradoAño string    `json:"especialidad_grado_año"`
	AñoDelDato           time.Time `json:"año_del_dato" `
	FechaDeGraduacion    time.Time `json:"fecha_de_graduacion" `
}

func (s *Server) UpdateEstudiosActuales(c *gin.Context) {
	var req UpdateEstudiosActuales
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.UpdateEstudiosActualesParams{
		TipoEnseñansa:         req.Tipo_enseñansa,
		Centro:                req.Centro,
		EspecialidadGradoOAño: req.EspecialidadGradoAño,
		AñoDelDato:            req.AñoDelDato,
		FechaDeGraduacion:     req.FechaDeGraduacion,
	}
	est, err := s.GetStore().UpdateEstudiosActuales(c, param)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "no estudios actuales found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating estudios actuales"})
		return
	}

	c.JSON(http.StatusOK, est)

}

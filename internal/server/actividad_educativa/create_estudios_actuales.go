package actividad_educativa

import (
	database "TesisAclifim/database/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateEstudiosActualesR struct {
	Tipo_enseñansa       string    `json:"tipo_ense" binding:"required"`
	Centro               string    `json:"centro" binding:"required"`
	EspecialidadGradoAño string    `json:"especialidad_grado_año" binding:"required"`
	AñoDelDato           time.Time `json:"año_del_dato" binding:"required"`
	FechaDeGraduacion    time.Time `json:"fecha_de_graduacion" binding:"required"`
}

func (s *Server) CreateEstudiosActuales(c *gin.Context) {
	var req CreateEstudiosActualesR
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.CreateEstudiosActualesParams{

		TipoEnseñansa:         req.Tipo_enseñansa,
		Centro:                req.Centro,
		EspecialidadGradoOAño: req.EspecialidadGradoAño,
		AñoDelDato:            req.AñoDelDato,
		FechaDeGraduacion:     req.FechaDeGraduacion,
	}
	est, err := s.GetStore().CreateEstudiosActuales(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, est)
}

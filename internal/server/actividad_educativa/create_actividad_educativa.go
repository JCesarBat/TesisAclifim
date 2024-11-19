package actividad_educativa

import (
	database "TesisAclifim/database/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActividadEducativaR struct {
	Id_asociado           int64  `json:"id_asociado" biding:"required"`
	IDEstudiosActuales    int64  `json:"id_estudios_actuales" biding:"required"`
	Ultimo_grado_aprobado string `json:"ultimo_grado_aprobado" biding:"required"`
}

func (s *Server) CreateActividadE(c *gin.Context) {
	var req ActividadEducativaR
	var UltimoG database.UltimoGrado
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := UltimoG.Scan(req.Ultimo_grado_aprobado)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := database.CreateActividadEducativaParams{
		IDAsociado:          req.Id_asociado,
		IDEstudiosActuales:  req.IDEstudiosActuales,
		UltimoGradoAprobado: UltimoG,
	}
	act, err := s.GetStore().CreateActividadEducativa(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, act)
}

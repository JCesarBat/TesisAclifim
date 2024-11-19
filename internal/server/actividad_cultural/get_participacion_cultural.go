package actividad_cultural

import (
	database "TesisAclifim/database/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetParticipacionCulturalRequest struct {
	Asiociado_Id int64 `uri:"asiociado_id"biding:"required"`
}

type GetParticipacionCulturalResponse struct {
	ID_asociado   int64                     `json:"id_asociado"`
	Aficcion      []string                  `json:"aficcion"`
	Participacion []database.ParticipacionC `json:"participacion"`
}

func (s *Server) GetParticipacionCultural(c *gin.Context) {
	var req GetParticipacionCulturalRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	act, err := s.GetStore().GetActividadCultural(c, req.Asiociado_Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	part, err := s.GetStore().GetParticipacionC(c, sql.NullInt64{Int64: act.ID, Valid: true})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetParticipacionCulturalResponse{
		ID_asociado:   act.ID,
		Aficcion:      act.Especialidad,
		Participacion: part,
	})
}

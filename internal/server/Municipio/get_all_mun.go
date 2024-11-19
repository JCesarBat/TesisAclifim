package Municipio

import (
	"TesisAclifim/internal/server/common_data"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllMunRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// @BasePath /municipio
// PingExample godoc
// @Summary 	get all mun
// @Schemes
// @Description  retorna todas los municipoios de una provincia
// @Tags Municipio
// @Produce json
// @Param register body GetAllMunRequest true "get mun"
// @Success 200 {object} database.Municipio
// @Router /municipio [get]
func (s *Server) GetAllMun(c *gin.Context) {
	var req GetAllMunRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}

	mun, err := s.GetStore().GetAllMunicipio(c, req.ID)
	if err != nil {
		c.JSON(500, common_data.ErrorResponse(err))
		return
	}
	c.JSON(200, mun)
	return
}

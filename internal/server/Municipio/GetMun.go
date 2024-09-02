package Municipio

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetMunRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// @BasePath /municipio
// PingExample godoc
// @Summary 	get a mun be a id
// @Schemes
// @Description  retorna un municipio segun el id
// @Tags Municipio
// @Accept json
// @Produce json
// @Param register body GetMunRequest true "get prov"
// @Success 200 {object} database.Municipio
// @Router /municipio [get]
func (s *Server) GetMun(c *gin.Context) {
	var req GetMunRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	mun, err := s.store.GetMunicipio(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, mun)
	return
}

package Prov

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetProvRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// @BasePath /provincias
// PingExample godoc
// @Summary 	get a prov be a id
// @Schemes
// @Description  retorna una provincia segun el id
// @Tags Provincia
// @Accept json
// @Produce json
// @Param register body GetProvRequest true "get prov"
// @Success 200 {object} database.Provincium
// @Router /provincia [get]
func (s *Server) GetProv(c *gin.Context) {
	var req GetProvRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	prov, err := s.store.GetProvincia(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, prov)
	return
}

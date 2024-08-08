package Prov

import "github.com/gin-gonic/gin"

// @BasePath /provincias
// PingExample godoc
// @Summary 	get all prov
// @Schemes
// @Description  retorna todas las provincias
// @Tags Provincia
// @Produce json
// @Success 200 {object} database.Provincium
// @Router /provincias [get]
func (s *Server) GetAllProv(c *gin.Context) {
	prov, err := s.store.GetAllProv(c)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, prov)
	return
}

package asociado

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) DinamicList(c *gin.Context) {
	var req database.DinamicListParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}

	resp, err := s.GetStore().DinamicList(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}

	c.JSON(200, resp)
}

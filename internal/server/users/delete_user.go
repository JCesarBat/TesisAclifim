package users

import (
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteUserRequest struct {
	Id int `uri:"id" binding:"required"`
}

// @BasePath /user/upgrade
// PingExample godoc
// @Summary 	delete a user by id
// @Schemes
// @Description  elimina un usuario segun su id
// @Tags users
// @Accept json
// @Produce json
// @Param  Update Password body	DeleteUserRequest true "delete user request"
// @Success 200 {string} success
// @Router /user/{id} [delete]
func (s *Server) DeleteUser(c *gin.Context) {
	var req DeleteUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	err := s.GetStore().DeleteUser(c, int64(req.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "success"})
}

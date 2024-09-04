package users

import (
	"Tesis/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}
type GetUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Municipio string    `json:"municipio"`
	Provincia string    `json:"provincia"`
	CreatedAt time.Time `json:"created_at"`
}

// @BasePath /user/{id}
// PingExample godoc
// @Summary 	get user by id
// @Schemes
// @Description  obtener un usuario por su id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} GetUserResponse
// @Router /user/{id} [get]
func (server *Server) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	user, err := server.GetStore().GetUserID(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}
	mun, prov, err := common_data.GetMunAndProv(user.IDMunicipio, server.GetStore(), c)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found the prov or the mun"})
			return
		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
	}

	response := GetUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		Municipio: mun.Name,
		Provincia: prov.Name,
		CreatedAt: user.CreatedAt,
	}
	c.JSON(http.StatusOK, response)
}

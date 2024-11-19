package users

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"TesisAclifim/pkg/util"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePasswordRequest struct {
	Id        int64  `json:"id" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Password2 string `json:"password2" binding:"required,min=8"`
}

// @BasePath /user/password
// PingExample godoc
// @Summary 	update a password
// @Schemes
// @Description  actualiza la contraseña de algun usuario
// @Tags users
// @Accept json
// @Produce json
// @Param  Update Password body	UpdatePasswordRequest true "update password request"
// @Success 200 {string} success
// @Router /user/password [put]
func (s *Server) UpdatePassword(c *gin.Context) {
	var req UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	if req.Password != req.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the passwords do not are equal"})
		return
	}
	HashPasswird, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}
	param := database.UpdateUserParams{
		ID:       req.Id,
		Password: sql.NullString{String: HashPasswird, Valid: true},
	}
	_, err = s.GetStore().UpdateUser(c, param)
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

type UpgradeToSuperUser struct {
	Id    int  `json:"id" binding:"required" `
	Valid bool `json:"valid,omitempty"`
}

// @BasePath /user/upgrade
// PingExample godoc
// @Summary 	upgrade to a super user
// @Schemes
// @Description  hace que un usuario se vuelva un super usuario
// @Tags users
// @Accept json
// @Produce json
// @Param  Update Password body	UpgradeToSuperUser true "update password request"
// @Success 200 {string} success
// @Router /user/upgrade [put]
func (s *Server) UpgradeToSuperUser(c *gin.Context) {
	var req UpgradeToSuperUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	user, err := s.GetStore().GetUserID(c, int64(req.Id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}
	if user.SuperUser.Bool == req.Valid {
		c.JSON(http.StatusPreconditionRequired, gin.H{"error": "the user is have the restriction selected"})
		return
	}
	param := database.UpdateToSuperUserParams{
		ID:        user.ID,
		SuperUser: sql.NullBool{Bool: req.Valid, Valid: true},
	}
	_, err = s.GetStore().UpdateToSuperUser(c, param)
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

package auth

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"TesisAclifim/pkg/util"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type LoginRequest struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	ID                   int64     `json:"id" `
	Username             string    `json:"Username"`
	Email                string    `json:"Email"`
	Municipio            string    `json:"municipio"`
	Provincia            string    `json:"provincia"`
	CreatedAt            time.Time `json:"created_at"`
	SessionID            uuid.UUID `json:"session_id"`
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpireAt  time.Time `json:"accessToken_expire_at"`
	RefreshToken         string    `json:"RefreshToken"`
	RefreshTokenExpireAt time.Time `json:"RefreshToken_expire_at"`
}

// @BasePath /login
// PingExample godoc
// @Summary 	login example
// @Schemes
// @Description  this is a example to login in the site
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login request"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func (s *Server) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("the cannot be null %v", err)})
		return
	}
	user, err := s.GetStore().GetUser(c, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("not found the user %v", err)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the password is invalid"})
		return
	}
	mun, err := s.GetStore().GetMunicipio(c, user.IDMunicipio)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("error in the database the municipio dont exists %v", err)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	prov, err := s.GetStore().GetProvincia(c, mun.IDProvincia)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("error in the database the provincia dont exists %v", err)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	accessToken, accessPayload, err := s.GetTokenMaker().CreateToken(
		user.Username,
		s.GetConfig().AccessTokenDuration,
	)

	if err != nil {
		c.JSON(http.StatusAlreadyReported, common_data.ErrorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := s.GetTokenMaker().CreateToken(
		user.Username,
		s.GetConfig().Refresh_Token_Duration,
	)

	if err != nil {
		c.JSON(http.StatusAlreadyReported, common_data.ErrorResponse(err))
		return
	}

	session, err := s.GetStore().CreateUSessions(c, database.CreateUSessionsParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    c.Request.UserAgent(),
		ClientIp:     c.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	response := LoginResponse{
		ID:                   user.ID,
		Username:             user.Username,
		Email:                user.Email,
		Municipio:            mun.Name,
		Provincia:            prov.Name,
		CreatedAt:            user.CreatedAt,
		SessionID:            session.ID,
		AccessToken:          accessToken,
		AccessTokenExpireAt:  accessPayload.ExpiredAt,
		RefreshToken:         refreshToken,
		RefreshTokenExpireAt: refreshPayload.ExpiredAt,
	}
	c.JSON(http.StatusOK, response)
}

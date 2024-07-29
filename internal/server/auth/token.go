package auth

import (
	"Tesis/internal/token"
	"Tesis/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type RefreshAccessTokenResponse struct {
	AccessToken         string    `json:"accessToken"`
	AccessTokenExpireAt time.Time `json:"accessToken_expire_at"`
}

// @BasePath /token
// PingExample godoc
// @Summary 	token auth
// @Schemes
// @Description  this is the documentation abaut token to refresh the token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body RefreshAccessTokenRequest true "refreshAccessTokenRequest"
// @Success 200 {object} RefreshAccessTokenResponse
// @Router /token [post]
// the refresh access token refresh the token passed and check if this is correct
func (server *Server) RefreshAccesToken(ctx *gin.Context) {
	var req RefreshAccessTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	config, err := util.LoadConfig("..")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetrickey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	refreshPayload, err := token.Maker.VerifyToken(tokenMaker, req.RefreshToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	session, err := server.store.GetSessions(ctx, refreshPayload.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("missmatch session token ")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusAlreadyReported, errorResponse(err))
		return
	}

	response := RefreshAccessTokenResponse{

		AccessToken:         accessToken,
		AccessTokenExpireAt: accessPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, response)

}

package auth

import (
	"TesisAclifim/internal/server/common_data"
	"TesisAclifim/internal/token"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPeyloadKey = "authorization_peyload"
)

// the authorization middleware check if the token is correct  or not
func AutMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		autorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		fields := strings.Fields(autorizationHeader)
		if len(autorizationHeader) == 0 {
			err := errors.New("authorization header is not provide")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common_data.ErrorResponse(err))
			return
		}

		if len(fields) < 2 {
			err := errors.New("Invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common_data.ErrorResponse(err))
			return
		}
		autorizationtype := strings.ToLower(fields[0])

		if autorizationtype != authorizationTypeBearer {
			err := errors.New("unsupported authorization")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common_data.ErrorResponse(err))
			return
		}

		accesToken := fields[1]
		peyload, err := tokenMaker.VerifyToken(accesToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common_data.ErrorResponse(err))
		}
		ctx.Set(authorizationHeaderKey, peyload)
		ctx.Next()
	}
}

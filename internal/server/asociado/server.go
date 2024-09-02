package asociado

import (
	database "Tesis/database/sqlc"
	"Tesis/internal/token"
	"Tesis/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      database.Store
	TokenMaker token.Maker
	config     util.Config
}

func NewServer(store database.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetrickey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	return &Server{
		store:      store,
		TokenMaker: tokenMaker,
		config:     config,
	}, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

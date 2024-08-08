package router

import (
	"Tesis/internal/server/auth"
	"Tesis/internal/server/users"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func InitRouter(authServer *auth.Server, userServer *users.Server) {
	r = gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// these are the routes of auth server
	r.POST("/login", authServer.Login)
	r.POST("/register", authServer.Register)
	r.POST("/token", authServer.RefreshAccesToken)
	authRouts := r.Group("/").Use(auth.AutMiddleware(authServer.TokenMaker))
	// users routes

	authRouts.GET("/user/:id", userServer.GetUser)
	authRouts.GET("/users", userServer.ListUser)
	authRouts.PUT("/user/password", userServer.UpdatePassword)
	authRouts.PUT("/user/upgrade", userServer.UpgradeToSuperUser)
	authRouts.DELETE("/user/:id", userServer.DeleteUser)

}

func Run(addr string) error {
	return r.Run(addr)
}
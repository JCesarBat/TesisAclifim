package router

import (
	"Tesis/internal/server/auth"
	"Tesis/internal/server/orchestrator"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func InitRouter(orchestrator *orchestrator.Orchestrator) {
	r = gin.Default()

	// Provincias y Municipios
	r.GET("/provincia", orchestrator.Prov.GetProv)
	r.GET("/provincias", orchestrator.Prov.GetAllProv)

	r.GET("/municipio", orchestrator.Mun.GetMun)
	r.GET("/municipios", orchestrator.Mun.GetAllMun)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// these are the routes of auth server
	r.POST("/login", orchestrator.Auth.Login)
	r.POST("/register", orchestrator.Auth.Register)
	r.POST("/token", orchestrator.Auth.RefreshAccesToken)
	authRouts := r.Group("/").Use(auth.AutMiddleware(orchestrator.Auth.GetTokenMaker()))
	// users routes

	authRouts.GET("/user/:id", orchestrator.User.GetUser)
	authRouts.GET("/users", orchestrator.User.ListUser)
	authRouts.PUT("/user/password", orchestrator.User.UpdatePassword)
	authRouts.PUT("/user/upgrade", orchestrator.User.UpgradeToSuperUser)
	authRouts.DELETE("/user/:id", orchestrator.User.DeleteUser)
	// asociado
	r.GET("/asociado/:id", orchestrator.Asociado.GetAsociado)
	r.GET("/asociados", orchestrator.Asociado.ListAsociados)
	r.POST("/asociado", orchestrator.Asociado.CreateAsociado)
	r.PUT("/asociado", orchestrator.Asociado.UpdateAsociado)
	r.DELETE("/asociado/:id", orchestrator.Asociado.DeleteAsociado)
}

func Run(addr string) error {
	return r.Run(addr)
}

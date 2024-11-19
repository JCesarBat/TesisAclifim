package router

import (
	"TesisAclifim/internal/server/orchestrator"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func InitRouter(orchestrator *orchestrator.Orchestrator) {
	r = gin.Default()

	// Provincias y Municipios
	r.GET("/provincia/:id", orchestrator.Prov.GetProv)
	r.GET("/provincias", orchestrator.Prov.GetAllProv)

	r.GET("/municipio/:id", orchestrator.Mun.GetMun)
	r.GET("/municipios/:id", orchestrator.Mun.GetAllMun)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// these are the routes of auth server
	r.POST("/login", orchestrator.Auth.Login)
	r.POST("/register", orchestrator.Auth.Register)
	r.POST("/token", orchestrator.Auth.RefreshAccesToken)
	//authRouts := r.Group("/").Use(auth.AutMiddleware(orchestrator.Auth.GetTokenMaker()))
	// users routes
	r.GET("/user/:id", orchestrator.User.GetUser)
	r.GET("/users", orchestrator.User.ListUser)
	r.PUT("/user/password", orchestrator.User.UpdatePassword)
	r.PUT("/user/upgrade", orchestrator.User.UpgradeToSuperUser)
	r.DELETE("/user/:id", orchestrator.User.DeleteUser)
	// asociado
	r.GET("/asociado/:id", orchestrator.Asociado.GetAsociado)
	r.GET("/asociados", orchestrator.Asociado.ListAsociados)
	r.GET("/asociado/dincamic", orchestrator.Asociado.DinamicList)
	r.POST("/asociado", orchestrator.Asociado.CreateAsociado)
	r.PUT("/asociado", orchestrator.Asociado.UpdateAsociado)
	r.DELETE("/asociado/:id", orchestrator.Asociado.DeleteAsociado)

	// Datos Sociales
	r.GET("/datosSocial/:id", orchestrator.DatosS.GetDatosS)
	r.POST("/datosSocial", orchestrator.DatosS.InsertDatosSociales)
	r.PUT("/datosSocial", orchestrator.DatosS.UpdateD)
	r.DELETE("/datosSocial", orchestrator.DatosS.DeleteDatosSociales)

	//Actividad Deportiva
	r.GET("/actividadDeportiva/:id", orchestrator.ActDepor.GetActividadDeportiva)
	r.POST("/actividadDeportiva", orchestrator.ActDepor.CreateAD)
	r.POST("/participacionDeportiva", orchestrator.ActDepor.CreateParticipacionD)
	r.PUT("/actividadDeportiva", orchestrator.ActDepor.UpdateActividadDeportiva)
	r.PUT("participacionDeportiva", orchestrator.ActDepor.UpdatePorticipacionDeportiva)
	r.DELETE("/actividadDeportiva/:id", orchestrator.ActDepor.DActividadDeportivaR)
	r.DELETE("/participacionDeportiva/:id", orchestrator.ActDepor.DeleteParticipacionDeportiva)

	// Actividad Educativa
	r.GET("ActividadEducativa/:id", orchestrator.ActEduc.GetActividadEducativa)
	r.POST("ActividadEducativa", orchestrator.ActEduc.CreateActividadE)
	r.POST("EstudiosActuales", orchestrator.ActEduc.CreateEstudiosActuales)
	r.PUT("ActividadEducativa", orchestrator.ActEduc.UpdateActividadEducativa)
	r.PUT("EstudiosActuales", orchestrator.ActEduc.UpdateEstudiosActuales)
	r.DELETE("ActividadEducativa/:id", orchestrator.ActEduc.DeleteActividadEducativa)
	r.DELETE("EstudiosActuales/:id", orchestrator.ActEduc.DeleteEstudiosActuales)

	// Actividad Cultural
	r.GET("ActividadCultural/:id", orchestrator.ActCult.GetParticipacionCultural)
	r.POST("ActividadCultural", orchestrator.ActCult.CreateActividadCultural)
	r.POST("ParticipacionCultural", orchestrator.ActCult.CreateParticipacion)
	r.PUT("actividadcultural", orchestrator.ActCult.UpdateActividadCultural)
	r.PUT("participacioncultural", orchestrator.ActCult.UpdateParticipacionCultural)
	r.DELETE("actividadcultural/:id", orchestrator.ActCult.DeleteActividadCultural)
	r.DELETE("participacioncultural/:id", orchestrator.ActCult.DeleteParticipacionCultural)

}

func Run(addr string) error {
	return r.Run(addr)
}

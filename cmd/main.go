package main

import (
	database "Tesis/database/sqlc"
	"Tesis/internal/server/Municipio"
	"Tesis/internal/server/Prov"
	"Tesis/internal/server/auth"
	"Tesis/internal/server/router"
	"Tesis/internal/server/users"
	_ "Tesis/pkg/docs"
	"Tesis/pkg/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a login API.
// @host localhost:8080
// @BasePath
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config")
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to load the database")
	}
	store := database.NewStore(db)
	authServer, err := auth.NewServer(store, config)
	if err != nil {
		log.Fatal("Failed to start the auth server")
	}
	userServer, err := users.NewServer(store, config)
	if err != nil {
		log.Fatal("Failed to start the users server")
	}
	provServer, err := Prov.NewServer(store, config)
	if err != nil {
		log.Fatal("Failed to start the prov server")
	}
	munServer, err := Municipio.NewServer(store, config)
	if err != nil {
		log.Fatal("Failed to start the mun server")
	}
	router.InitRouter(authServer, userServer, provServer, munServer)
	log.Println("the server is running in the port:", config.HTTP_Server)
	router.Run(config.HTTP_Server)
}

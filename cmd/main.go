package main

import (
	database "Tesis/database/sqlc"
	"Tesis/internal/server/orchestrator"
	"Tesis/internal/server/router"
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
	orchestatrator, err := orchestrator.NewOrchestrator(store, config)
	if err != nil {
		log.Fatal("faild to load all servers", err.Error())
	}
	router.InitRouter(orchestatrator)
	log.Println("the server is running in the port:", config.HTTP_Server)
	router.Run(config.HTTP_Server)
}

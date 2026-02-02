package main

import (
	"fmt"
	"log"
	"user-auth-service/config"
	"user-auth-service/db"
	"user-auth-service/handler"
	"user-auth-service/repo"
	"user-auth-service/router"
	"user-auth-service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	config, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load environmet variables: %v", err)
	}

	db, err := db.Connect(config)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := repo.NewRepo(db, logger)
	service := service.NewService(repo, logger)
	handler := handler.NewHandler(service, logger)

	r := gin.Default()

	router.SetUpRoutes(r, handler)

	r.Run(fmt.Sprintf(":%s", config.SERVER_ADDRESS))
}

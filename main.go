package main

import (
	"fmt"
	"log"
	"user-auth-service/config"
	"user-auth-service/db"
	user_repo "user-auth-service/repo/user"
	user_service "user-auth-service/service/user"
	user_handler "user-auth-service/handler/user"
	"user-auth-service/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load environmet variables: %v", err)
	}

	db, err := db.Connect(config)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := user_repo.NewRepo(db, logger)
	service := user_service.NewService(repo, logger)
	handler := user_handler.NewHandler(service, logger)

	r := gin.Default()

	router.SetUpRoutes(r, handler)

	r.Run(fmt.Sprintf(":%s", config.SERVER_ADDRESS))
}

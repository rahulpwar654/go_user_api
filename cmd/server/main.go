// @title Gin User API
// @version 1.0
// @description Simple user API with GORM + MySQL
// @host localhost:8080
// @BasePath /
//
//go:generate swag init -g main.go -o ../../docs -parseDependency -parseInternal
package main

import (
	"log"

	_ "gin-user-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-user-api/internal/controller"
	"gin-user-api/internal/db"
	"gin-user-api/internal/model"
	"gin-user-api/internal/repository"
	"gin-user-api/internal/service"
)

func main() {
	r := gin.Default() // logger + recovery middleware
	// serve swagger JSON at /doc.json and swagger UI at /swagger/*any
	// prefer generated swagger.json, fallback to manual file
	r.StaticFile("/doc.json", "docs/swagger.json")
	r.StaticFile("/doc.manual.json", "docs/swagger_manual.json")
	// default the swagger UI to the manual spec so UI shows API immediately
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/doc.manual.json")))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	cfg := db.LoadConfig()
	dbConn, err := db.OpenMySQL(cfg)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	if err := dbConn.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("db migrate failed: %v", err)
	}

	repo := repository.NewGormUserRepository(dbConn)
	svc := service.NewUserService(repo)
	uc := controller.NewUserController(svc)
	uc.RegisterRoutes(r)

	log.Println("starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

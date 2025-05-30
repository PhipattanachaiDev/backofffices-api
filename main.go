package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	database "ezview.asia/ezview-web/ezview-lite-back-office/configs/databases"
	docs "ezview.asia/ezview-web/ezview-lite-back-office/docs"
	router "ezview.asia/ezview-web/ezview-lite-back-office/routers"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// @title eZView Lite Back Office API
// @version 1.0
// @description For eZVeiw Lite Back Office Project.
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

type Config struct {
	GinMode  string
	Port     string
	BasePath string
}

func loadConfig() (*Config, error) {
	// ใช้ godotenv เฉพาะในระหว่างการพัฒนา
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			return nil, errors.Wrap(err, "Error loading .env file")
		}
	}

	config := &Config{
		GinMode:  os.Getenv("GIN_MODE"),
		Port:     os.Getenv("PORT"),
		BasePath: os.Getenv("BASE_PATH"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.BasePath == "" {
		config.BasePath = "/back-office-service/api"
	}

	return config, nil
}

func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.BasePath = config.BasePath + "/api"

	// Set Gin mode based on environment variable
	if config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to the database
	db := database.ConnectDB()
	defer db.Close()

	// Setup router
	r := router.SetupRouter()

	// Server setup
	srv := &http.Server{
		Addr: "0.0.0.0:" + config.Port,
		// Addr:         "127.0.0.1:" + config.Port,
		Handler:      r,
		ReadTimeout:  120 * time.Second, // ตั้งค่าเวลารอการอ่าน
		WriteTimeout: 120 * time.Second, // ตั้งค่าเวลารอการเขียน
		IdleTimeout:  120 * time.Second, // ตั้งค่าเวลาที่เซิร์ฟเวอร์จะรอการเชื่อมต่อที่ว่าง
	}

	// Log message for starting server
	log.Printf("Starting server on port %s...", config.Port)

	// Log message for Swagger UI
	log.Printf("Swagger UI available at http://localhost:%s%s/swagger/index.html", config.Port, config.BasePath)

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

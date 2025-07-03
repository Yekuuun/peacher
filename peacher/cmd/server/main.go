package main

import (
	"fmt"
	"log"
	"os"
	"peacher/internal/builder"
	"peacher/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[!] ERROR : .env file not found.")
		return
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Println("[!] ERROR : Telegram bot token not found.")
		return
	}

	config := builder.NewAppBuilder().
		WithPort(8080).
		WithTelegramToken(token).
		Build()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // use server IP ?
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(server, *config)

	if run := server.Run(fmt.Sprintf(":%d", config.Port)); run != nil {
		log.Printf("[!] ERROR : Error loading app on port %d \n", config.Port)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"peacher/internal/builder"
	"peacher/routes"

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

	chatId := os.Getenv("BOT_CHAT_ID")
	if chatId == "" {
		log.Println("[!] ERROR : Telegram bot token not found.")
		return
	}

	config := builder.NewAppBuilder().
		WithPort(8080).
		WithTelegramToken(token).
		WithChatId(chatId).
		Build()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	routes.RegisterRoutes(server, *config)

	if run := server.Run(fmt.Sprintf(":%d", config.Port)); run != nil {
		log.Printf("[!] ERROR : Error loading app on port %d \n", config.Port)
	}
}

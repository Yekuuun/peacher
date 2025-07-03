package routes

import (
	"peacher/internal/builder"
	"peacher/internal/notify"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, config builder.App) {
	handler := notify.NewHandler(config.TelegramToken, config.ChatId)

	router.GET("/status", handler.GetStatus)
	router.POST("/payload", handler.SendPayload)
}

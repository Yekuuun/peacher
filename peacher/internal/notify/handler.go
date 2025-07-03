package notify

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Token string
}

func NewHandler(token string) *Handler {
	return &Handler{
		Token: token,
	}
}

func (h *Handler) SendPayload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (h *Handler) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

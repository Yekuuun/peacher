package notify

type MaliciousPayload struct {
	Identifier string `json:"identifier"` //username || email.
	Password   string `json:"password"`
}

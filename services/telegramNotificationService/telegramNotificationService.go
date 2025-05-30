package telegramNotificationService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Initialize environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

const telegramAPI = "https://api.telegram.org/bot"

// SendNotification sends a message to a Telegram bot.
func SendNotification(message string) error {
	if message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	// Get the bot token and chat ID from environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if botToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is missing")
	}
	if chatID == "" {
		return fmt.Errorf("TELEGRAM_CHAT_ID environment variable is missing")
	}

	enabledStr := os.Getenv("TELEGRAM_NOTIFY_ENABLED")
	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		return fmt.Errorf("error parsing TELEGRAM_NOTIFY_ENABLED: %w", err)
	}

	if !enabled {
		fmt.Println("Telegram notifications are disabled. Notification not sent.")
		return nil
	}

	// Telegram API URL
	apiURL := fmt.Sprintf("%s%s/sendMessage", telegramAPI, botToken)

	// Create the message payload
	payload := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Create a POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body for additional error details
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification. Status code: %d, Response: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

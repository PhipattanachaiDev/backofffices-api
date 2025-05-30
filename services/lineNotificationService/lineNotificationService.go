package lineNotificationService

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Initialize environment variables
func init() {
	if os.Getenv("GIN_MODE") != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v\n", err)
		}
	}
}

const lineNotifyAPI = "https://notify-api.line.me/api/notify"

// SendNotification sends a message to Line Notify.
func SendNotification(message string) error {
	if message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	token := os.Getenv("LINE_NOTIFY_TOKEN")
	if token == "" {
		return fmt.Errorf("LINE_NOTIFY_TOKEN environment variable is missing")
	}

	enabledStr := os.Getenv("LINE_NOTIFY_ENABLED")
	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		return fmt.Errorf("error parsing LINE_NOTIFY_ENABLED: %w", err)
	}

	if !enabled {
		fmt.Println("Line Notify is disabled. Notification not sent.")
		return nil
	}

	// Create request body
	formData := url.Values{}
	formData.Set("message", message)

	req, err := http.NewRequest("POST", lineNotifyAPI, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return fmt.Errorf("error creating new request: %w", err)
	}

	// Set authorization header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenForAuth string

func TestLoginV1(t *testing.T) {
	r := SetupRouter()

	// Perform a POST request to /v1/login
	loginData := map[string]string{
		"username": "foo",
		"password": "bar",
	}
	jsonData, _ := json.Marshal(loginData)

	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert response status code and capture token
	assert.Equal(t, http.StatusOK, w.Code)

	var loginResponse map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)

	token, ok := loginResponse["data"].(map[string]interface{})["token"].(string)
	assert.True(t, ok)

	// Store token for subsequent tests or pass it to other test functions
	tokenForAuth = token
}

func TestPingV1(t *testing.T) {
	r := SetupRouter()

	// Prepare a GET request to /v1/ping with authorization header
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	req.Header.Set("Authorization", "Bearer "+tokenForAuth) // Use the token obtained from TestLogin

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert response status code
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), `"success":true`)
	assert.Contains(t, w.Body.String(), `"status":200`)
	assert.Contains(t, w.Body.String(), `"message":"OK"`)
}

func TestPingV2(t *testing.T) {
	r := SetupRouter()

	// Prepare a GET request to /v1/ping with authorization header
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	req.Header.Set("Authorization", "Bearer "+tokenForAuth) // Use the token obtained from TestLogin

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert response status code
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), `"success":true`)
	assert.Contains(t, w.Body.String(), `"status":200`)
	assert.Contains(t, w.Body.String(), `"message":"OK"`)
}

func TestPingV3(t *testing.T) {
	r := SetupRouter()

	// Prepare a GET request to /v1/ping with authorization header
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	req.Header.Set("Authorization", "Bearer "+tokenForAuth) // Use the token obtained from TestLogin

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert response status code
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), `"success":true`)
	assert.Contains(t, w.Body.String(), `"status":200`)
	assert.Contains(t, w.Body.String(), `"message":"OK"`)
}

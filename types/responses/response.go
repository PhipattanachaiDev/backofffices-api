package responses

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response struct defines the structure of API responses
type Response struct {
	Success   bool        `json:"success" extensions:"x-order=0"`
	Status    int         `json:"status" extensions:"x-order=1"`
	Timestamp time.Time   `json:"timestamp" extensions:"x-order=2"`
	Message   string      `json:"message" extensions:"x-order=3"`
	Data      interface{} `json:"data,omitempty" extensions:"x-order=4"`
}

// NewResponse creates a new response with the specified status code, message, and optional data
func NewResponse(success bool, status int, message string, data interface{}) Response {
	return Response{
		Success:   success,
		Status:    status,
		Timestamp: time.Now().UTC(),
		Message:   message,
		Data:      data,
	}
}

// JSON formats the response as JSON and writes it to the gin context
func (resp Response) JSON(c *gin.Context) {
	resp.Timestamp = resp.Timestamp.UTC()
	c.JSON(resp.Status, resp)
}

// ErrorJSON formats an error response with an error message
func ErrorJSON(c *gin.Context, status int, err error) {
	response := Response{
		Success:   false,
		Status:    status,
		Timestamp: time.Now().UTC(),
		Message:   err.Error(),
	}
	c.JSON(status, response)
}

// OK sends a 200 OK response with optional data
func OK(c *gin.Context, data interface{}) {
	response := NewResponse(true, http.StatusOK, "OK", data)
	response.JSON(c)
}

// Created sends a 201 Created response without data
func Created(c *gin.Context) {
	response := NewResponse(true, http.StatusCreated, "Created", nil)
	response.JSON(c)
}

// CreatedData sends a 201 Created response with data
func CreatedData(c *gin.Context, data interface{}) {
	response := NewResponse(true, http.StatusCreated, "Created", data)
	response.JSON(c)
}

// Updated sends a 200 OK response with a message indicating successful update without data
func Updated(c *gin.Context) {
	response := NewResponse(true, http.StatusOK, "Updated", nil)
	response.JSON(c)
}

// NotFoundWithData sends a 200 OK response with data indicating not found condition
func NotFoundWithData(c *gin.Context, data interface{}) {
	response := NewResponse(true, http.StatusOK, "Not Found", data)
	c.JSON(http.StatusOK, response)
}

// BadRequest sends a 400 Bad Request response with an error message
func BadRequest(c *gin.Context, message string) {
	response := Response{
		Success:   false,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now().UTC(),
		Message:   message,
	}
	c.JSON(http.StatusBadRequest, response)
}

// Unauthorized sends a 401 Unauthorized response with an error message
func Unauthorized(c *gin.Context, message string) {
	response := Response{
		Success:   false,
		Status:    http.StatusUnauthorized,
		Timestamp: time.Now().UTC(),
		Message:   message,
	}
	c.JSON(http.StatusUnauthorized, response)
}

// Forbidden sends a 403 Forbidden response with an error message
func Forbidden(c *gin.Context, message string) {
	response := Response{
		Success:   false,
		Status:    http.StatusForbidden,
		Timestamp: time.Now().UTC(),
		Message:   message,
	}
	c.JSON(http.StatusForbidden, response)
}

// NotFound sends a 404 Not Found response with an error message
func NotFound(c *gin.Context, message string) {
	response := Response{
		Success:   false,
		Status:    http.StatusNotFound,
		Timestamp: time.Now().UTC(),
		Message:   message,
	}
	c.JSON(http.StatusNotFound, response)
}

// InternalServerError sends a 500 Internal Server Error response with an error message
func InternalServerError(c *gin.Context, message string) {
	response := Response{
		Success:   false,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now().UTC(),
		Message:   message,
	}
	c.JSON(http.StatusInternalServerError, response)
}

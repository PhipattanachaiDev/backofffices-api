package certificatesController

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreatePDFHandler(t *testing.T) {
	// สร้าง router
	router := gin.Default()
	router.GET("/api/v1/certificates/pdf", CreatePDF)

	// สร้าง request
	req, err := http.NewRequest("GET", "/api/v1/certificates/pdf", nil)
	if err != nil {
		t.Fatal(err)
	}

	// สร้าง response recorder
	w := httptest.NewRecorder()

	// รัน request
	router.ServeHTTP(w, req)

	// ตรวจสอบ status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", w.Code)
	}

	// ตรวจสอบ Content-Type
	if w.Header().Get("Content-Type") != "application/pdf" {
		t.Errorf("Expected Content-Type 'application/pdf', got %v", w.Header().Get("Content-Type"))
	}
}

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"websiteapi/middleware"
	"websiteapi/service"

	"github.com/gin-gonic/gin"
)

func TestAuthorizedMiddleware(t *testing.T) {
	// Initialize a new Gin router
	r := gin.Default()

	// Add your authorization middleware to the router
	r.Use(middleware.Authorized())

	// Create a test route that requires authorization
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You are authorized!",
		})
	})

	// Create a test request with a valid JWT token in the Authorization header
	req := httptest.NewRequest("GET", "/protected", nil)
	token, _ := service.CreateToken(123) // Replace 123 with a valid user ID
	authHeader := "Bearer " + token
	req.Header.Set("Authorization", authHeader)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, w.Code)
	}

	// Optionally, you can check the response body or other assertions
	// For example, you can assert that the response body contains a specific message

	expectedResponse := `{"message":"You are authorized!"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, w.Body.String())
	}
}

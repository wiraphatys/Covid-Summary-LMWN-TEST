package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouter set up the Gin router for testing
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/covid/summary", covidSummaryHandler)
	return r
}

// TestCovidSummaryEndpoint test the /covid/summary endpoint
func TestCovidSummaryEndpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check if the response contain the correct key
	assert.Contains(t, response, "Province")
	assert.Contains(t, response, "AgeGroup")

	// Check the structure of the "AgeGroup" key
	ageGroup, ok := response["AgeGroup"].(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, ageGroup, "0-30")
	assert.Contains(t, ageGroup, "31-60")
	assert.Contains(t, ageGroup, "61+")
	assert.Contains(t, ageGroup, "N/A")
}

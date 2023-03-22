package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAlbumsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/albums", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		getAlbums(c)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var albums []album
	if err := json.Unmarshal(rr.Body.Bytes(), &albums); err != nil {
		t.Errorf("error unmarshaling response body: %v", err)
	}

}

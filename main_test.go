package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestPostAlbumsHandler(t *testing.T) {
	// create a new album to add
	newAlbum := album{
		ID:     "3",
		Title:  "New Album",
		Artist: "New Artist",
		Price:  9.99,
	}

	// marshal the new album into JSON format
	newAlbumJSON, err := json.Marshal(newAlbum)
	if err != nil {
		t.Fatal(err)
	}

	// create a new request with the JSON payload
	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(newAlbumJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a new recorder to record the response
	rr := httptest.NewRecorder()

	// create a new router and add the postAlbums handler
	r := gin.Default()
	r.POST("/albums", postAlbums)

	// call the postAlbums handler with the new request
	r.ServeHTTP(rr, req)

	// check if the response status code is as expected
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// check if the response body contains the new album
	var responseAlbum album
	err = json.Unmarshal(rr.Body.Bytes(), &responseAlbum)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}
	if responseAlbum != newAlbum {
		t.Errorf("Handler returned unexpected body: got %v want %v", responseAlbum, newAlbum)
	}

	// Check if the new album exists in the slice
	found := false
	for _, a := range albums {
		if a == newAlbum {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected slice to contain new album %v, but it doesn't exist in the slice.", newAlbum)
	}
}

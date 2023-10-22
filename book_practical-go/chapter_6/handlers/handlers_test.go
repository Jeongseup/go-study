package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/config"
)

func TestApiHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()

	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	apiHandler(w, r, c)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// t.Log(resp)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected response status: %v, Got: %v\n", http.StatusOK, resp.StatusCode)
	}

	expectedResponseBody := "Hello, world!"

	if string(body) != expectedResponseBody {
		t.Errorf("Expected response: %s, Got: %s\n", expectedResponseBody, string(body))
	}
}
func TestHealthHandler(t *testing.T) {
	r := httptest.NewRequest("POST", "/healthz", nil)
	w := httptest.NewRecorder()

	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	healthCheckHandler(w, r, c)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// t.Log(resp)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected response status: %v, Got: %v\n", http.StatusOK, resp.StatusCode)
	}

	expectedResponseBody := "Hello, world!"

	if string(body) != expectedResponseBody {
		t.Errorf("Expected response: %s, Got: %s\n", expectedResponseBody, string(body))
	}
}

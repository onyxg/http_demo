package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDelayImageMiddleware(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/images/test.jpg", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := delayImageMiddleware(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	start := time.Now()
	if err := h(c); err != nil {
		t.Fatalf("handler returned an error: %v", err)
	}
	duration := time.Since(start)

	if duration < 2*time.Second {
		t.Errorf("expected delay of at least 2 seconds, got %v", duration)
	}
}

func TestMainFunction(t *testing.T) {
	// This test will check if the main function runs without errors
	go func() {
		main()
	}()

	// Give the server some time to start
	time.Sleep(1 * time.Second)

	// Check if the server is running on port 9001
	resp, err := http.Get("http://localhost:9001")
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", resp.Status)
	}
}

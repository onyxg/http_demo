package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetDummyJSON(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/dummy-json", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getDummyJSON(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := `{"name":"John Doe","age":30,"sex":"male","friends":["Alice","Bob","Charlie"]}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}

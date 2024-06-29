package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetWord(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/word", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	expected := "word"

	if err := getWord(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %d", rec.Code)
	}

	if rec.Body.String() != expected {
		t.Errorf("expected body %q; got %q", expected, rec.Body.String())
	}
}

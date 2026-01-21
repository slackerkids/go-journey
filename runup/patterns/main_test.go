package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	if Sum(2, 3) != 5 {
		t.Fatal("expected 5")
	}
}

func TestIsEven(t *testing.T) {
	cases := []struct {
		input int
		want  bool
	}{
		{2, true},
		{3, false},
	}

	for _, c := range cases {
		if IsEven(c.input) != c.want {
			t.Errorf("input %d", c.input)
		}
	}
}

func TestHandler(t *testing.T) {
	body := strings.NewReader(`{"name":"World"}`)
	req := httptest.NewRequest(http.MethodPost, "/hello", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
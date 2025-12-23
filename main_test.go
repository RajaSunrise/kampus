package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	// Initialize templates
	loadTemplates()

	// Setup router
	mux := setupRouter()

	tests := []struct {
		route string
	}{
		{"/"},
		{"/about"},
		{"/academics"},
		{"/admissions"},
		{"/research"},
		{"/student-life"},
		{"/news"},
		{"/directory"},
		{"/contact"},
	}

	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	loadTemplates()
	mux := setupRouter()

	req, err := http.NewRequest("GET", "/non-existent-page", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

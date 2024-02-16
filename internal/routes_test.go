package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {

	// Setup
	Get()

	t.Run("Get  StatusOK", func(t *testing.T) {
		// Given
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		rec := httptest.NewRecorder()

		// When
		http.DefaultServeMux.ServeHTTP(rec, req)

		// Then
		if rec.Result().StatusCode != http.StatusOK {
			t.Fatal("StatusCode ", rec.Result().StatusCode)
		}
		if rec.Body.String() != "testing get" {
			t.Fatal("Body: ", rec.Body)
		}
	})

	t.Run("Head  StatusOK", func(t *testing.T) {
		// Given
		req := httptest.NewRequest(http.MethodHead, "/test", nil)
		rec := httptest.NewRecorder()

		// When
		http.DefaultServeMux.ServeHTTP(rec, req)

		// Then
		if rec.Result().StatusCode != http.StatusOK {
			t.Fatal("StatusCode ", rec.Result().StatusCode)
		}
		if rec.Body.String() != "testing get" {
			t.Fatal("Body: ", rec.Body)
		}
	})

	t.Run("Post  StatusMethodNotAllowed", func(t *testing.T) {
		// Given
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		rec := httptest.NewRecorder()

		// When
		http.DefaultServeMux.ServeHTTP(rec, req)

		// Then
		if rec.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Fatal("StatusCode ", rec.Result().StatusCode)
		}
		if rec.Body.String() != "Method Not Allowed\n" {
			t.Fatal("Body: ", rec.Body)
		}
	})

}

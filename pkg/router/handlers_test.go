package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	t.Run("Hitting the Health Check", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		res := httptest.NewRecorder()
		NewRouter().ServeHTTP(res, req)
		assert.Equal(t, 200, res.Code, "OK response is expected")
	})
}

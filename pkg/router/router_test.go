package router

import "testing"

func TestNewRouter(t *testing.T) {
	t.Run("If router has no error", func(t *testing.T) {
		got := NewRouter()
		if got == nil {
			t.Log("No error starting the router")
		}
	})
}

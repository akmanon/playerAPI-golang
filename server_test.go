package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Return Rahul score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Rahul", nil)
		res := httptest.NewRecorder()

		PlayerServer(res, req)
		got := res.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRecordWinsAndRetrieveThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Sonu"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinReq(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinReq(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinReq(player))

	res := httptest.NewRecorder()

	server.ServeHTTP(res, newGetScoreReq(player))
	assert.Equal(t, res.Code, http.StatusOK)
	assert.Equal(t, res.Body.String(), "3")
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	mu.Lock()
	defer mu.Unlock()
	s.winCalls = append(s.winCalls, name)
	s.scores[name]++
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{map[string]int{
		"Rahul": 20,
		"Karan": 30,
	},
		nil,
	}

	server := &PlayerServer{&store}
	t.Run("return Rahul score", func(t *testing.T) {
		req := newGetScoreReq("Rahul")
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, res.Body.String(), "20")

	})

	t.Run("return Karan score", func(t *testing.T) {
		req := newGetScoreReq("Karan")
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, res.Body.String(), "30")

	})
	t.Run("return 404 on missing players", func(t *testing.T) {
		req := newGetScoreReq("")
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, res.Code, http.StatusNotFound)

	})
}

func newGetScoreReq(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it records win when post and return accepted on POST", func(t *testing.T) {
		req := NewPostWinReq("Rahul")
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, res.Code, http.StatusAccepted)
		assert.Len(t, store.winCalls, 1)
	})
}

func NewPostWinReq(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return req
}

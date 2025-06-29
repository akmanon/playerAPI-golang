package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
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

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{map[string]int{
		"Rahul": 20,
		"Karan": 30,
	},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)
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
		nil,
	}
	server := NewPlayerServer(&store)

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

func TestLeague(t *testing.T) {

	t.Run("it returns League table as Json", func(t *testing.T) {
		wantedLeague := []Player{
			{"Karan", 20},
			{"Rahul", 30},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		req := NewLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		got := GetLeagueFromResponse(t, res.Body)

		gotContentType := res.Result().Header.Get("content-type")
		jsonContentType := "application/json"

		assert.Equal(t, gotContentType, jsonContentType, "got %q, want %q", gotContentType, jsonContentType)
		assert.Equal(t, res.Code, http.StatusOK, "Expected Status Code to be equal")
		assert.Equal(t, got, wantedLeague, "got %v, want %v", got, wantedLeague)
	})
}

func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league/", nil)
	return req
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	assert.NoError(t, err, "Unable to Parse %+v", league)
	return
}

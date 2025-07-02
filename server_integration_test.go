package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRecordWinsAndRetrieveThem(t *testing.T) {
	db, cleandb := createTempFile(t, "[]")
	defer cleandb()
	store, err := NewFsPlayerStore(db)
	server := NewPlayerServer(store)
	player := "Sonu"
	assert.NoError(t, err)
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinReq(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinReq(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinReq(player))
	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreReq(player))
		assert.Equal(t, res.Code, http.StatusOK)
		assert.Equal(t, res.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Sonu", 3},
		}
		res := httptest.NewRecorder()
		req := NewLeagueRequest()

		server.ServeHTTP(res, req)

		got := GetLeagueFromResponse(t, res.Body)

		assert.Equal(t, res.Code, http.StatusOK, "Expected Status Code to be equal")
		assert.Equal(t, got, wantedLeague, "got %v, want %v", got, wantedLeague)

	})

}

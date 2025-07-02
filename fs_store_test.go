package poker

import (
	"os"
	"testing"

	"github.com/alecthomas/assert"
)

func TestFsStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		db, cleandb := createTempFile(t, `[
			{"Name": "Karan", "Wins" : 399},
			{"Name": "Rahul", "Wins" : 34}]`)
		defer cleandb()
		store, err := NewFsPlayerStore(db)
		got := store.GetLeague()
		want := League{
			{"Karan", 399},
			{"Rahul", 34},
		}
		assert.NoError(t, err)
		assert.Equal(t, got, want, "got %v, want %v", got, want)
	})
	t.Run("Get Player score", func(t *testing.T) {
		db, cleandb := createTempFile(t, `[
			{"Name": "Karan", "Wins" : 3},
			{"Name": "Rahul", "Wins" : 34}]`)
		defer cleandb()
		store, err := NewFsPlayerStore(db)
		got := store.GetPlayerScore("Karan")
		want := 3
		assert.NoError(t, err)
		assert.Equal(t, got, want, "got %v, want %v", got, want)

	})
	t.Run("record win for existing player", func(t *testing.T) {
		db, cleandb := createTempFile(t, `[
			{"Name": "Karan", "Wins" : 3},
			{"Name": "Rahul", "Wins" : 34}]`)
		defer cleandb()
		store, err := NewFsPlayerStore(db)
		store.RecordWin("Karan")
		got := store.GetPlayerScore("Karan")
		want := 4
		assert.NoError(t, err)
		assert.Equal(t, got, want, "got %v, want %v", got, want)

	})
	t.Run("record win for new player", func(t *testing.T) {
		db, cleandb := createTempFile(t, `[
			{"Name": "Karan", "Wins" : 3},
			{"Name": "Rahul", "Wins" : 34}]`)
		defer cleandb()
		store, err := NewFsPlayerStore(db)
		store.RecordWin("Sonu")
		got := store.GetPlayerScore("Sonu")
		want := 1
		assert.NoError(t, err)
		assert.Equal(t, got, want, "got %v, want %v", got, want)

	})
	t.Run("Work with an empty file", func(t *testing.T) {
		db, cleandb := createTempFile(t, "")
		defer cleandb()
		_, err := NewFsPlayerStore(db)
		assert.NoError(t, err)

	})
	//file_system_store_test.go
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFsPlayerStore(database)

		assert.NoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.Equal(t, got, want)

		// read again
		got = store.GetLeague()
		assert.Equal(t, got, want)
	})

}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Errorf("err while creating temp file, %v", err)
	}
	tmpfile.WriteString(initialData)
	removeTmp := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeTmp
}

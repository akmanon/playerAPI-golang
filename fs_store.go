package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FsPlayerStore struct {
	db     *json.Encoder
	league League
}

func FsPlayerStoreFromFile(path string) (*FsPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s, %v", path, err)
	}
	closeFunc := func() {
		db.Close()
	}
	store, err := NewFsPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating fs player store %v", err)
	}
	return store, closeFunc, nil
}

func NewFsPlayerStore(file *os.File) (*FsPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store file %s, %v ", file.Name(), err)
	}
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store file %s, %v ", file.Name(), err)
	}

	return &FsPlayerStore{
		db:     json.NewEncoder(&tape{file}),
		league: league,
	}, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v ", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func (f *FsPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FsPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FsPlayerStore) RecordWin(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.db.Encode(f.league)
}

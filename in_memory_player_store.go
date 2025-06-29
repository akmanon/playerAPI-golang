package main

import "sync"

type InMemoryPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

var mu sync.Mutex

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	mu.Lock()
	defer mu.Unlock()
	s.winCalls = append(s.winCalls, name)
	s.scores[name]++
}

// in_memory_player_store.go
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.scores {
		league = append(league, Player{name, wins})
	}
	return league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, []string{}, []Player{}}
}

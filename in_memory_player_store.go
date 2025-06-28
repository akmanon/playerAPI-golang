package main

import "sync"

type InMemoryPlayerStore struct {
	scores   map[string]int
	winCalls []string
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

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, []string{}}
}

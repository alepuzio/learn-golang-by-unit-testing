package main

import "sync"

/*
NewInMemoryPlayerStore initialises an empty player store.
*/
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

/*
InMemoryPlayerStore collects data about players in memory.
*/
type InMemoryPlayerStore struct {
	store map[string]int
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex //occupies the resources
}

/*
RecordWin will record a player's win.
*/
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()//occupies the resources
	defer i.lock.Unlock()//free resources
	i.store[name]++
}

/*
GetPlayerScore retrieves scores for a given player. Lock the resources
*/
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock() //occupies the resources
	defer i.lock.RUnlock() //free resources
	return i.store[name]
}

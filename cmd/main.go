package main

import (
	"github.com/zlw8844/playerscore/server"
	"log"
	"net/http"
)

// main.go

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	players map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.players[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.players[name]
}

func main() {
	server := server.NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5002", server))
}

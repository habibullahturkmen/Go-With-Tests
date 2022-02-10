package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
		if name == "Pepper" {
			return  20
		}

		if name == "Floyd" {
			return 10
		}
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5050", server))
}

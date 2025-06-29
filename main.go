package main

import (
	"log"
	"net/http"
)

func main() {
	playerServer := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", playerServer))
}

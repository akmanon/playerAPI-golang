// main.go
package main

import (
	"log"
	"net/http"

	poker "github.com/akmanon/playerAPI-golang"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FsPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

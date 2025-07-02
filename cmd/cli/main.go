package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/akmanon/playerAPI-golang"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")
	store, close, err := poker.FsPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	game := poker.NewCli(store, os.Stdin)
	game.PlayPoker()

}

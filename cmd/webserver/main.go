package main

import (
	. "github.com/zlw8844/playerscore"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, _ := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5002", server))
}

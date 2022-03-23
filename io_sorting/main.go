package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

// The 2nd argument to os.OpenFile lets you define the permissions for opening the file, in our case O_RDWR means we want to read and write and os.O_CREATE means create the file if it doesn't exist.
// The 3rd argument means sets permissions for the file, in our case, all users can read and write the file
func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

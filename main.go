package main

import (
	"log"
	"net/http"
)

func main() {
	database, cleanDatabaes := NewFileSystemPlayerStore()
	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}

package main

import (
	"log"
	"net/http"
	"trackserver/server"
	"trackserver/store"
)

func main() {
	store := store.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}

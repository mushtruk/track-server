package main

import (
	"log"
	"net/http"
	"trackserver/server"
	"trackserver/store"
)

func main() {
	store := store.NewInMemoryPlayerStore()
	server := &server.PlayerServer{}
	server.SetStore(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}

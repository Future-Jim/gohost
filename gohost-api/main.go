package main

import (
	"log"

	"github.com/future-jim/gohost/lib/userstore"
)

func main() {
	store, err := userstore.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()

}

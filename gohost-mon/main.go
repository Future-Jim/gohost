package main

import (
	"github.com/future-jim/gohost/lib/storage"
)

const measurementDelay = 1

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		return
	}
	store.Init()
	agent := NewMetricAgent(store)
	agent.Run()
}

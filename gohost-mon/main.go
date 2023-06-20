package main

import (
	"log"

	"github.com/future-jim/gohost/lib/metricstore"
)

const measurementDelay = 60

func main() {
	store, err := metricstore.NewPostgresStore()
	if err != nil {
		return
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	agent := NewMetricAgent(store)
	agent.Run()
}

package main

import "github.com/future-jim/gohost/lib/metricstore"

const measurementDelay = 1

func main() {
	store, err := metricstore.NewPostgresStore()
	if err != nil {
		return
	}
	store.Init()
	agent := NewMetricAgent(store)
	agent.Run()
}

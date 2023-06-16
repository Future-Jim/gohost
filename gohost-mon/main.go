package main

import (
	"fmt"
	"time"

	"github.com/future-jim/gohost/lib/storage"
	"github.com/future-jim/gohost/lib/types"
)

const measurementDelay = 60

func main() {
	var metric types.Metrics
	store, err := storage.NewPostgresStore()
	if err != nil {
		return
	}
	store.Init()
	for {

		time.Sleep(measurementDelay * time.Second)
		PMU := GetMemoryPercentUsed()
		AL := GetAverageLoad()
		HUT := GetHostUpTime()

		metric.AL = AL
		metric.PMU = PMU
		metric.HUT = HUT
		store.AddEntry(&metric)
		fmt.Println(metric)
	}
}

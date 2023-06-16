package main

import (
	"math"
	"time"

	"github.com/future-jim/gohost/lib/storage"
	"github.com/future-jim/gohost/lib/types"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

const bytesToMB = 1048576
const secondsInHour = 3600
const secondsInDay = 86400
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
		PMU := getMemoryPercentUsed()
		AL := getAverageLoad()
		HUT := getHostUpTime()

		metric.AL = AL
		metric.PMU = PMU
		metric.HUT = HUT
		store.AddEntry(&metric)

	}
}

func getMemoryPercentUsed() types.PercentMemoryUsed {
	v, _ := mem.VirtualMemory()
	return types.PercentMemoryUsed{
		PMU: int(math.Trunc(v.UsedPercent)),
	}
}
func getHostUpTime() types.HostUpTime {
	h, _ := host.Info()
	days, _ := divmod(h.Uptime, secondsInDay)
	hours, minutes := divmod(h.Uptime, secondsInHour)
	return types.HostUpTime{
		Days:    days,
		Hours:   hours,
		Minutes: minutes / 60,
	}
}

func getAverageLoad() types.AverageLoad {
	l, _ := load.Avg()
	var l1 float64 = roundFloat(l.Load1, 2)
	var l5 float64 = roundFloat(l.Load5, 2)
	var l15 float64 = roundFloat(l.Load15, 2)

	return types.AverageLoad{
		One:     l1,
		Five:    l5,
		Fifteen: l15,
	}
}

func divmod(numerator, denominator uint64) (quotient, remainder uint64) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

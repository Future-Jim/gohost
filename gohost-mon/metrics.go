package main

import (
	"fmt"
	"math"
	"time"

	"github.com/future-jim/gohost/lib/metricstore"
	"github.com/future-jim/gohost/lib/types"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

const bytesToMB = 1048576
const secondsInHour = 3600
const secondsInDay = 86400

type MetricAgent struct {
	name  string
	store metricstore.MetricStorage
}

func NewMetricAgent(store metricstore.MetricStorage) *MetricAgent {
	return &MetricAgent{
		store: store,
	}
}

func (a *MetricAgent) Run() {
	var metric types.Metrics
	for {
		time.Sleep(measurementDelay * time.Second)
		metric.AL = GetAverageLoad()
		metric.PMU = GetMemoryPercentUsed()
		metric.HUT = GetHostUpTime()
		a.store.AddEntry(&metric)
		fmt.Println(metric)
		result := a.store.GetEntry()
		fmt.Println(result)
	}
}

func GetMemoryPercentUsed() types.PercentMemoryUsed {
	v, _ := mem.VirtualMemory()
	return types.PercentMemoryUsed{
		PMU: int(math.Trunc(v.UsedPercent)),
	}
}
func GetHostUpTime() types.HostUpTime {
	h, _ := host.Info()
	days, _ := divMod(h.Uptime, secondsInDay)
	hours, minutes := divMod(h.Uptime, secondsInHour)
	return types.HostUpTime{
		Days:    days,
		Hours:   hours,
		Minutes: minutes / 60,
	}
}

func GetAverageLoad() types.AverageLoad {
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

func divMod(numerator, denominator uint64) (quotient, remainder uint64) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

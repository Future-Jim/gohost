package main

import (
	"fmt"
	"time"

	storage "github.com/future-jim/gohost-storage"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

func main() {
	var metric Metrics
	storage.NewPostgresStore()

	for {

		time.Sleep(1 * time.Second)
		PMU := getMemoryPercentUsed()
		AL := getAverageLoad()
		HUT := getHostUpTime()

		metric.AL = AL
		metric.PMU = PMU
		metric.HUT = HUT
		fmt.Println(metric)

	}
}

func getMemoryPercentUsed() PercentMemoryUsed {
	v, _ := mem.VirtualMemory()
	return PercentMemoryUsed{
		PMU: v.UsedPercent,
	}
}
func getHostUpTime() HostUpTime {
	h, _ := host.Info()
	days, _ := divmod(h.Uptime, secondsInDay)
	hours, minutes := divmod(h.Uptime, secondsInHour)
	return HostUpTime{
		Days:    days,
		Hours:   hours,
		Minutes: minutes / 60,
	}
}

func getAverageLoad() AverageLoad {
	l, _ := load.Avg()
	return AverageLoad{
		One:     l.Load1,
		Five:    l.Load5,
		Fifteen: l.Load15,
	}
}

func divmod(numerator, denominator uint64) (quotient, remainder uint64) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return quotient, remainder
}

package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

const bytesToMB = 1048576
const secondsInHour = 3600
const secondsInDay = 86400

type HostUpTime struct {
	Days    uint64
	Hours   uint64
	Minutes uint64
}

type AverageLoad struct {
	One     float64
	Five    float64
	Fifteen float64
}

func main() {

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(getMemoryPercentUsed())
		fmt.Println(getHostUpTime())
		fmt.Println(getAverageLoad())
	}
}

func getMemoryPercentUsed() float64 {
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
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

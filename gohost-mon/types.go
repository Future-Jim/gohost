package main

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

type PercentMemoryUsed struct {
	PMU float64
}

type Metrics struct {
	HUT HostUpTime
	AL  AverageLoad
	PMU PercentMemoryUsed
}

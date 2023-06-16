package types

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
	PMU int
}

type Metrics struct {
	HUT HostUpTime
	AL  AverageLoad
	PMU PercentMemoryUsed
}

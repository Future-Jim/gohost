module github.com/future-jim/gohost

replace github.com/future-jim/gohost/lib/storage => ../lib/storage
replace github.com/future-jim/gohost/lib/types => ../lib/types

go 1.20

require (
	github.com/future-jim/gohost/lib/storage v0.0.0-00010101000000-000000000000
	github.com/future-jim/gohost/lib/types v1.0.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/shirou/gopsutil/v3 v3.23.5
)

require (
	github.com/future-jim/gohost/lib/types v1.0.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

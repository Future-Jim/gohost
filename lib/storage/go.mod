module github.com/future-jim/gohost/lib/storage

replace github.com/future-jim/gohost/lib/types => ../types

go 1.20

require (
	github.com/lib/pq v1.10.9 // indirect
	github.com/future-jim/gohost/lib/types v1.0.0
)
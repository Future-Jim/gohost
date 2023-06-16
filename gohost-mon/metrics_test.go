package main

import (
	"reflect"
	"testing"

	"github.com/future-jim/gohost/lib/types"
)

func TestGetMemoryPercentUsed(t *testing.T) {
	got := GetMemoryPercentUsed()
	want := new(types.PercentMemoryUsed)

	if reflect.TypeOf(&got) != reflect.TypeOf(want) {
		t.Error("did not return a struct of type", reflect.TypeOf(want))
	}
}

func TestGetAverageLoad(t *testing.T) {
	got := GetAverageLoad()
	want := new(types.AverageLoad)
	if reflect.TypeOf(&got) != reflect.TypeOf(want) {
		t.Error("did not return a struct of type", reflect.TypeOf(want))
	}
}

func TestGetHostUpTime(t *testing.T) {
	got := GetHostUpTime()
	want := new(types.HostUpTime)
	if reflect.TypeOf(&got) != reflect.TypeOf(want) {
		t.Error("did not return a struct of type", reflect.TypeOf(want))
	}
}

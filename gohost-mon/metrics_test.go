package main

import (
	"reflect"
	"testing"
)

func TestGetMemoryPercentUsed(t *testing.T) {
	got := GetMemoryPercentUsed()
	//dumby int to test against return type in GetMemoryPercentUsed
	var want int
	if reflect.TypeOf(got.PMU) != reflect.TypeOf(want) {
		t.Error("Did not return a value of type", reflect.TypeOf(want))
	}
}

func TestGetAverageLoad(t *testing.T) {
	got := GetAverageLoad()
	var want float64
	if reflect.TypeOf(got.One) != reflect.TypeOf(want) {
		t.Error("load_1 returned incorrect type => ", reflect.TypeOf(got.One))
	}
	if reflect.TypeOf(got.Five) != reflect.TypeOf(want) {
		t.Error("load_1 returned incorrect type =>", reflect.TypeOf(want))
	}
	if reflect.TypeOf(got.Fifteen) != reflect.TypeOf(want) {
		t.Error("load_1 returned incorrect type =>", reflect.TypeOf(want))
	}
}

func TestGetHostUpTime(t *testing.T) {
	got := GetHostUpTime()
	var want uint64
	if reflect.TypeOf(got.Days) != reflect.TypeOf(want) {
		t.Error("GetHostUpTime returned incorrect type for day:", reflect.TypeOf(got.Days))
	}
	if reflect.TypeOf(got.Hours) != reflect.TypeOf(want) {
		t.Error("GetHostUpTime returned incorrect type for day:", reflect.TypeOf(got.Hours))
	}
	if reflect.TypeOf(got.Minutes) != reflect.TypeOf(want) {
		t.Error("GetHostUpTime returned incorrect type for day:", reflect.TypeOf(got.Minutes))
	}
}

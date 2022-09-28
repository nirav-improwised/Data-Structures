package methods_test

import (
	"main/methods"
	"testing"
)

func TestCuisine(t *testing.T) {
	records := methods.GetRecords()
	got := methods.Cuisine(0, records, "Japanese")
	expect := 1
	if got != int8(expect) {
		t.Errorf("Expected %v got %v", expect, got)
	}
}

func TestCity(t *testing.T) {
	records := methods.GetRecords()
	got := methods.City(0, records, "Makati City")
	expect := 1
	if got != int8(expect) {
		t.Errorf("Expected %v got %v", expect, got)
	}
}

func TestHasTableBooking(t *testing.T) {
	records := methods.GetRecords()
	got := methods.HasTableBooking(0, records, false)
	expect := 1
	if got != int8(expect) {
		t.Errorf("Expected %v got %v", expect, got)
	}
}

package methods_test

import (
	"main/methods"
	"testing"
)

// Restaurant ID,Restaurant Name,Country Code,City,Address,Locality,Locality Verbose,Longitude,Latitude,Cuisines,Average Cost for two,Currency,Has Table booking,Has Online delivery,Is delivering now,Switch to order menu,Price range,Aggregate rating,Rating color,Rating text,Votes

var records = methods.GetRecords()
var dataStructs = methods.GeneratingStructs(records)

func TestCuisine(t *testing.T) {
	got := methods.Cuisine(0, dataStructs, "Japanese")
	expect := 1
	if got != int8(expect) {
		t.Errorf("Expected %v got %v", expect, got)
	}
}

func TestCity(t *testing.T) {
	got := methods.City(0, dataStructs, "Makati City")
	expect := 1
	if got != int8(expect) {
		t.Errorf("Expected %v got %v", expect, got)
	}
}

func TestHasTableBooking(t *testing.T) {
	got := methods.HasTableBooking(0, dataStructs, "No")
	expect := 1
	if got != int8(expect) {
		t.Errorf("Expected %v got %v", expect, got)
	}
}

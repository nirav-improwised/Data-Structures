package methods

import (
	"encoding/csv"
	"fmt"
	"strings"

	"main/structures"
	"os"
)

func GeneratingStructs(records [][]string) []structures.Data {
	// const length = len(records)
	// var data = [length]structures.Data{}

	// data := make([]structures.Data, len(records))
	data := []structures.Data{}

	for i := range records {

		// s := structures.Data{records[i][0], records[i][1], records[i][2], records[i][3], records[i][4], records[i][5], records[i][6], records[i][7], records[i][8], records[i][9], records[i][10], records[i][11], records[i][12], records[i][13], records[i][14], records[i][15], records[i][16], records[i][17], records[i][18], records[i][19], records[i][20]}
		s := structures.Data{
			RestaurantID:      records[i][0],
			RestaurantName:    records[i][1],
			CountryCode:       records[i][2],
			City:              records[i][3],
			Address:           records[i][4],
			Locality:          records[i][5],
			Verbose:           records[i][6],
			Longitude:         records[i][7],
			Latitude:          records[i][8],
			Cuisines:          records[i][9],
			AverageCost:       records[i][10],
			Currency:          records[i][11],
			HasTableBooking:   records[i][12],
			HasOnlineDelivery: records[i][13],
			IsDeliveringNow:   records[i][14],
			SwitchToOrderMenu: records[i][15],
			PriceRange:        records[i][16],
			AggregateRating:   records[i][17],
			RatingColor:       records[i][18],
			RatingText:        records[i][19],
			Votes:             records[i][20],
		}
		data = append(data, s)
	}
	return data
}

func GetRecords() [][]string {
	fp, err := os.Open("/home/nirav/workspace/src/DataStrcutures/Zomato_CLI_problem/zomato.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	fileReader := csv.NewReader(fp)
	Records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return Records
}
func Cuisine(i int, data []structures.Data, cuisine string) int8 {
	var flag int8 = 1
	if cuisine != "default" {
		if !(strings.Contains(data[i].Cuisines, cuisine)) {
			flag = 0
		}
	}
	return flag
}

func City(i int, data []structures.Data, city string) int8 {
	var flag int8 = 1
	if city != "default" {
		if !(strings.Contains(data[i].City, city)) {
			flag = 0
		}
	}
	return flag
}

func HasTableBooking(i int, data []structures.Data, hTB bool) int8 {
	var flag int8 = 1
	if hTB {
		if !(strings.Contains(data[i].HasTableBooking, "Yes")) {
			flag = 0
		}
	}
	return flag
}

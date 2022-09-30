package methods

import (
	"encoding/csv"
	"fmt"
	"strings"

	"main/structures"
	"os"
)

func GeneratingStructs(records [][]string) []structures.Data {

	data := []structures.Data{}

	for i := range records {

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

func HasTableBooking(i int, data []structures.Data, hTB string) int8 {
	var flag int8 = 1
	if hTB == "Yes" {
		if !(strings.Contains(data[i].HasTableBooking, "Yes")) {
			flag = 0
		}
	}
	return flag
}

func StudyData(dataStructs []structures.Data, cuisine, city, hTB string) []structures.TwoFieldData {
	count := 0
	var returnData []structures.TwoFieldData
	var temp structures.TwoFieldData
	for i := range dataStructs {
		var flag1, flag2, flag3 int8 = 1, 1, 1

		flag1 = Cuisine(i, dataStructs, cuisine)
		flag2 = City(i, dataStructs, city)
		flag3 = HasTableBooking(i, dataStructs, hTB)

		if flag1*flag2*flag3 == 1 {
			count++
			temp.Sr_No = count
			temp.RestaurantName = dataStructs[i].RestaurantName
			temp.Address = dataStructs[i].Address
			returnData = append(returnData, temp)
		}
	}
	return returnData
}

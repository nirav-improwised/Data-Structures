package methods

import (
	"encoding/csv"
	"fmt"
	"reflect"
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

func ConvertedData(newR structures.Data) [][]string {
	v := reflect.ValueOf(newR)
	temp := make([]string, v.NumField())
	var values [][]string
	for i := 0; i < v.NumField(); i++ {
		x := v.Field(i)
		temp[i] = x.String()
	}
	values = append(values, temp)
	return (values)
}

func AddData(mainStr [][]string) string {
	fp, err := os.OpenFile("/home/nirav/workspace/src/DataStrcutures/Zomato_CLI_problem/zomato.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	w := csv.NewWriter(fp)
	w.UseCRLF = true
	w.WriteAll(mainStr)

	return "Data Added"
}

func DeleteData(id string, records [][]string) string {
	flag := 0
	message := ""
	e := os.Remove("zomato.csv")
	if e != nil {
		fmt.Println(e)
	}
	fp, err := os.Create("zomato.csv")
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(fp)
	for i := range records {
		if id != records[i][0] {
			w.Write(records[i])
		} else {
			flag = 1
		}
	}
	w.Flush()
	fp.Close()

	if flag == 1 {
		message = fmt.Sprintf("%v Record Deleted Succuessfully", id)
	} else {
		message = "Record not found, kindly enter appropriate ID"
	}
	return message
}

func UpdateData(queryData structures.Data, records [][]string) string {

	for i := range records {
		if queryData.RestaurantID == records[i][0] {
			v := reflect.ValueOf(queryData)
			for j := 0; j < v.NumField(); j++ {
				x := v.Field(j)
				if !x.IsZero() {
					records[i][j] = x.String()
				}
			}
			break
		}
	}
	e := os.Remove("zomato.csv")
	if e != nil {
		fmt.Println(e)
	}
	fp, err := os.Create("zomato.csv")
	if err != nil {
		fmt.Println(err)
	}
	w := csv.NewWriter(fp)
	w.WriteAll(records)
	return fmt.Sprintf("%v Data Updated successfully", queryData.RestaurantID)
}

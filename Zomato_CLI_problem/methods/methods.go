package methods

import (
	"encoding/csv"
	"fmt"
	"main/zomato_cli_problem/cmd"
	"os"
	"strings"
)

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
func Cuisine(i int, records [][]string, cuisine string) int8 {
	var flag int8 = 1
	if cuisine != "default" {
		if !(strings.Contains(records[i][9], cmd.Cuisine)) {
			flag = 0
		}
	}
	return flag
}

func City(i int, records [][]string, city string) int8 {
	var flag int8 = 1
	if city != "default" {
		if !(strings.Contains(records[i][3], cmd.City)) {
			flag = 0
		}
	}
	return flag
}

func HasTableBooking(i int, records [][]string, hTB bool) int8 {
	var flag int8 = 1
	if hTB {
		if !strings.Contains(records[i][12], "Yes") {
			flag = 0
		}
	}
	return flag
}

package main

import (
	"encoding/csv"
	"fmt"
	"main/zomato_cli_problem/cmd"
	"os"
	"strings"
)

func main() {

	cmd.Execute()

	fp, err := os.Open("zomato.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	fileReader := csv.NewReader(fp)
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	flag1 := 1
	flag2 := 1
	flag3 := 1
	count := 0

	for i := range records {
		flag1, flag2, flag3 = 1, 1, 1

		if cmd.Cuisine != "default" {
			if !(strings.Contains(records[i][9], cmd.Cuisine)) {
				flag1 = 0
				continue
			}
		}
		if cmd.City != "default" {
			if !(strings.Contains(records[i][3], cmd.City)) {
				flag2 = 0
				continue
			}
		}
		if cmd.HasTableBooking {
			if !strings.Contains(records[i][12], "Yes") {
				flag3 = 0
				continue
			}
		}
		if flag1*flag2*flag3 == 1 {
			count++
			fmt.Println(count, records[i][1])
		}
	}
}

package main

import (
	"fmt"
	"main/methods"
	"main/zomato_cli_problem/cmd"
)

func main() {

	cmd.Execute()

	var flag1 int8
	var flag2 int8
	var flag3 int8
	count := 0
	Records := methods.GetRecords()
	for i := range Records {
		flag1, flag2, flag3 = 1, 1, 1
		cuisine := cmd.Cuisine
		city := cmd.City
		hTB := cmd.HasTableBooking

		flag1 = methods.Cuisine(i, Records, cuisine)
		flag2 = methods.City(i, Records, city)
		flag3 = methods.HasTableBooking(i, Records, hTB)
		if flag1*flag2*flag3 == 1 {
			count++
			fmt.Println(count, Records[i][1])
		}
	}
}

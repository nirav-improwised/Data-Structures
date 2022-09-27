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
	for i := range records {
		if strings.Contains(records[i][9], cmd.Cuisine) && strings.Contains(records[i][3], cmd.City) {
			fmt.Println(records[i][1])
		}
	}
}

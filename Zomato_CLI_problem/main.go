package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	var cuisine string
	fp, err := os.Open("zomato.csv")
	fmt.Println("Enter a cuise please: ")
	fmt.Scan(&cuisine)
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
		if strings.Contains(records[i][9], cuisine) {
			fmt.Println(records[i][1])
		}
	}
}

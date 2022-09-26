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
	cuisine = os.Args[2]
	city := strings.Join(os.Args[3:], " ")
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
		if strings.Contains(records[i][9], cuisine) && strings.Contains(records[i][3], city) {
			fmt.Println(records[i][1])
		}
	}
}

package day2

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readCsvFile(filePath string) []string {
	// Load a csv file.
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error trying to open file -", err)
	}
	// Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			return record
		}

		if err != nil {
			panic(err)
		}
	}
}

func parseData(stringData []string) []int {
	var intData []int
	for i, val := range stringData {
		intData[i], _ = strconv.Atoi(val)
	}
	return intData
}

func main() {
	stringData := readCsvFile("day-2-input.csv")
	intData := parseData(stringData)
}
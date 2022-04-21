package main

import (
	"fmt"
	"os"

	"encoding/csv"

	"github.com/NuVeS/CSVParser/finder"
	"github.com/NuVeS/CSVParser/parser"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter command: ")
	// comand, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Printf("Error %e", err)
	// 	return
	// }

	comand := "age = 27"

	expr, err := parser.Parse(comand)
	if err != nil {
		return
	}

	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = 3
	if err != nil {
		panic(err)
	}

	csvData, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	rows := finder.Find(csvData, expr)

	fmt.Println(rows)
}

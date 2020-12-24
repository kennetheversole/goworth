package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func csvWriter(date string, assets string, debts string, delta string) {

	// Open the file
	recordFile, err := os.OpenFile("finances.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	// Initialize the writer
	writer := csv.NewWriter(recordFile)

	// Create data to write
	var csvData = [][]string{
		{date, assets, debts, delta},
	}

	err = writer.WriteAll(csvData) // returns error
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

}

func main() {
	fmt.Println("Network Tracker")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter  today's date [mm/dd/yyyy]: ")
	scanner.Scan()
	date := scanner.Text()
	fmt.Print("Please enter asset value as of that date: ")
	scanner.Scan()
	assets := scanner.Text()
	fmt.Print("Please enter debt value as of that date: ")
	scanner.Scan()
	debts := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	assetsFLT64, err := strconv.ParseFloat(assets, 64)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	debtFLT64, err := strconv.ParseFloat(debts, 64)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	deltaFLT64 := assetsFLT64 - debtFLT64
	delta := fmt.Sprintf("%f", deltaFLT64)

	csvWriter(date, assets, debts, delta)

}

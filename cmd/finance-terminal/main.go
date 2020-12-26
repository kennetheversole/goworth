package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func writeToCSV(date string, assets string, debts string, delta string) {

	// Open the file
	recordFile, err := os.OpenFile("./finances.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Initialize the writer
	writer := csv.NewWriter(recordFile)

	// Create data to write
	var csvData = [][]string{
		{date, assets, debts, delta},
	}

	err = writer.WriteAll(csvData)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func main() {
	fmt.Println("===============")
	fmt.Println("Networth Tracker")
	fmt.Println("===============")

	now := time.Now()
	today := fmt.Sprintf("%02d/%02d/%d",
		now.Day(), now.Month(), now.Year())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nPlease enter asset value as of %v: ", today)
	scanner.Scan()
	assets := scanner.Text()
	fmt.Printf("\nPlease enter debt value as of %v: ", today)
	scanner.Scan()
	debts := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	assetsFLT64, err := strconv.ParseFloat(assets, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	debtFLT64, err := strconv.ParseFloat(debts, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	delta := fmt.Sprintf("%.2f", assetsFLT64-debtFLT64)

	writeToCSV(today, assets, debts, delta)

}

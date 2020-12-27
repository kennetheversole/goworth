package backend

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func WriteToCSV(date string, assets string, debts string, delta string) {

	// Open the file
	recordFile, err := os.OpenFile("./frontend/finances.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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

func Terminal(assets string, debts string) {

	today := time.Now().Format("02/01/2006")
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

	WriteToCSV(today, assets, debts, delta)

}

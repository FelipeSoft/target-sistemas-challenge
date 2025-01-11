package exercise

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Invoice struct {
	DateOfTheMonth int     `json:"dia"`
	Value          float64 `json:"valor"`
}

func RunExercise03() {
	jsonFilePath := "./../../assets/invoices.json"

	invoices, err := parseJSONToInvoices(jsonFilePath)
	if err != nil {
		log.Fatalf("Error on JSON file processing: %v", err)
	}

	minValue, maxValue, daysAboveAverage := calculateSalesMetrics(invoices)

	fmt.Printf("The minimum invoice value: %.2f\n", minValue)
	fmt.Printf("The maximum invoice value: %.2f\n", maxValue)
	fmt.Printf("Number of days with sales above the average: %d\n", daysAboveAverage)
}

func parseJSONToInvoices(filePath string) ([]Invoice, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var invoices []Invoice
	if err := json.Unmarshal(bytes, &invoices); err != nil {
		return nil, err
	}

	return invoices, nil
}

func calculateSalesMetrics(invoices []Invoice) (float64, float64, int) {
	var total, min, max float64
	var daysAboveAverage int
	var validValues []float64

	if len(invoices) > 0 {
		min = invoices[0].Value
		max = invoices[0].Value
	}

	for _, invoice := range invoices {
		value := invoice.Value

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}

		total += value
		validValues = append(validValues, value)
	}

	average := total / float64(len(validValues))

	for _, value := range validValues {
		if value > average {
			daysAboveAverage++
		}
	}

	return min, max, daysAboveAverage
}

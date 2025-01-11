package exercise

import "fmt"

func RunExercise04() {
	revenue := map[string]float64{
		"SP":    67836.43,
		"RJ":    36678.66,
		"MG":    29229.88,
		"ES":    27165.48,
		"Others": 19849.53,
	}

	var totalRevenue float64
	for _, value := range revenue {
		totalRevenue += value
	}

	for state, value := range revenue {
		percentage := (value / totalRevenue) * 100
		fmt.Printf("Percentage of revenue from %s: %.2f%%\n", state, percentage)
	}
}

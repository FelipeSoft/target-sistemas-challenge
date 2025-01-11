package exercise

import "fmt"

func RunExercise02() {
	var number int
	fmt.Print("Provide a number to verify if it belongs to the Fibonacci sequence: ")
	fmt.Scan(&number)

	if belongsToFibonacci(number) {
		fmt.Printf("The number %d belongs to the Fibonacci sequence.\n", number)
	} else {
		fmt.Printf("The number %d does not belong to the Fibonacci sequence.\n", number)
	}
}

func belongsToFibonacci(number int) bool {
	a, b := 0, 1

	if number == a || number == b {
		return true
	}

	for b <= number {
		a, b = b, a+b
		if b == number {
			return true
		}
	}

	return false
}

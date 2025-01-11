package exercise

import "fmt"

func RunExercise01() {
	index := 13
	sum := 0

	for k := 0; k < index; k++ {
		sum += k
	}

	fmt.Printf("Current Sum: %d \n", sum);
}
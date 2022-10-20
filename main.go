package main

import "fmt"

const OvenTime = 40

func main() {
	remaining := RemainingOvenTime(30)
	fmt.Println("Remaining oven time:", remaining)

	preparationTime := PreparationTime(2)
	fmt.Println("Preparation time: ", preparationTime)

	elapsedTime := ElapsedTime(3, 20)
	fmt.Println("Elapsed time: ", elapsedTime)
}

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

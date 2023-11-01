package main

import (
	"fmt"
)

func main() {
	var weight float64
	var height float64

	fmt.Print("Enter your weight in kilograms: ")
	fmt.Scanln(&weight)

	fmt.Print("Enter your height in meters: ")
	fmt.Scanln(&height)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI is %.2f\n", bmi)
}

// https://github.com/PinkOrca-net/Go-Fun

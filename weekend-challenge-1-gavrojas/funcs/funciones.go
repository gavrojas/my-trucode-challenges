package funcs

import (
	"fmt"
	"math"
	"strings"
)

func Temperature(temp int64, system string) string {
	// declaramos variables
	var message string
	var result int64

	switch strings.ToUpper(system) {
	case "F":
		// Cálculo de celsius a fahrenheit
		result = (temp * 9 / 5) + 32
		message = fmt.Sprintf("%d Celsius equals %d Fahrenheit", temp, result)

		return message

	case "C":
		// Cálculo de fahrenheit a celsius
		result = (temp - 32) * 5 / 9
		message = fmt.Sprintf("%d Fahrenheit equals %d Celsius", temp, result)

		return message

	default:
		message = fmt.Sprintln("You have to use 'C' or 'F'")
		return message
	}
}

func Bmi(weigth, heigth float64) (string, string) {
	// declaramos variables
	var bmiResult, message string

	result := weigth / math.Pow(heigth, 2)
	bmiResult = fmt.Sprintf("Right now your BMI is %.2f", result)

	if result < 18.5 {
		message = fmt.Sprintln("You are underweight, add more potato to the broth")
	} else if result >= 18.5 && result < 25 {
		message = fmt.Sprintln("You have a normal weight, I have healthy envy of you")
	} else {
		message = fmt.Sprintln("You are overweight, I know, the pandemic has affected us all")
	}

	return bmiResult, message
}

func Mario(heigth int) string {
	// declaramos variable
	var message string

	for i := 1; i <= heigth; i++ {
		for spaces := 0; spaces < heigth-i; spaces++ {
			fmt.Print(" ")
		}

		for numerals := 0; numerals < i; numerals++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	return message
}

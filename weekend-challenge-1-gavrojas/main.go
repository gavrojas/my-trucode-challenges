package main

import (
	"fmt"
	"log"
	"os"

	"trucode.app/w1challenge/funcs"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No elegiste ninguna opción \n\nEl uso correcto es alguna de las siguientes 2 opciones:\n./w1challenge [temperature|bmi|mario] \n./go run . [temperature|bmi|mario]")
	}

	switch action := os.Args[1]; action {
	case "temperature":
		// declaramos variables
		var temp int64
		var system string

		// Preguntamos datos a persona usuaria
		fmt.Println("What's the temperature?")
		fmt.Scanln(&temp)

		fmt.Println("To which system?(C/F)")
		fmt.Scanln(&system)

		//Imprimir resultado
		fmt.Println(funcs.Temperature(temp, system))

	case "bmi":
		// declaramos variables
		var weigth float64
		var heigth float64

		// Preguntamos datos a persona usuaria
		fmt.Println("How much do you weigh [kilograms]? (don't lie)")
		fmt.Scanln(&weigth)

		fmt.Println("How tall are you [meters]? (barefoot)")
		fmt.Scanln(&heigth)

		//Imprimir resultado
		bmi, result := funcs.Bmi(weigth, heigth)
		fmt.Printf("%s \n%s", bmi, result)

	case "mario":
		// declaramos variable
		var heigth int

		// Validamos rango entre 1 y 8
		for {
			fmt.Print("Pyramid height: ")
			fmt.Scanln(&heigth)
			if heigth >= 1 && heigth <= 8 {
				break
			}
		}
		//Imprimir resultado
		fmt.Print(funcs.Mario(heigth))
	default:
		fmt.Println("Operación no disponible")
		// podría haber usado log.fatal también
	}
}

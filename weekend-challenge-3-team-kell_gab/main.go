package main

import (
	"fmt"

	"trucode.com/interactions/food"
	"trucode.com/interactions/humans"
)

type humanAction int

const (
	UNSELECTED humanAction = iota
	EXERCISE
	SLEEP
	STUDY
	EAT
	SHOWSTATUS
	EXIT
)

type personOption int

const (
	HUMAN_UNSELECTED personOption = iota
	PERSON
	CHILD
)

// type foodType int

// const (
// 	UNSELECTED_FOOD foodType = iota
// 	DESSERT
// 	MEAL
// 	ENERGIZER
// )

func main() {

	var humanType personOption
	var personSelected humans.Humans
	var humanAction humanAction

	// Primer menú
	for humanType == HUMAN_UNSELECTED {

		fmt.Println("What did you want to create?:")
		fmt.Println("1. Person")
		fmt.Println("2. Child")

		_, err := fmt.Scanln(&humanType)
		if err != nil {
			humanType = HUMAN_UNSELECTED
		}

		switch humanType {
		case PERSON:
			person := humans.NewPerson("Wences")
			fmt.Println(person.ShowStatus())
			personSelected = person
		case CHILD:
			child := humans.NewChild("Wences Jr")
			fmt.Println(child.ShowStatus())
			personSelected = child
		default:
			humanType = HUMAN_UNSELECTED
			fmt.Println("Option not available")
		}

	}

	for humanAction != EXIT {
		fmt.Println("What did you want me to do?")
		fmt.Println("1. Exercise")
		fmt.Println("2. Sleep")
		fmt.Println("3. Study")
		fmt.Println("4. Eat")
		fmt.Println("5. Show Status")
		fmt.Println("6. Exit")

		_, err := fmt.Scan(&humanAction)
		if err != nil {
			humanAction = UNSELECTED
		}

		switch humanAction {
		case EXERCISE:
			intensity := humans.NO_INTENSITY

			for intensity == humans.NO_INTENSITY {
				var intensityOption int
				fmt.Println("Let me know with which intensity I should exercise:")
				fmt.Println("1. Low")
				fmt.Println("2. Medium")
				fmt.Println("3. High")

				_, err := fmt.Scan(&intensityOption)
				if err != nil {
					intensityOption = 0
				}

				switch intensityOption {
				case 1:
					intensity = humans.LOW_INTENSITY
				case 2:
					intensity = humans.MEDIUM_INTENSITY
				case 3:
					intensity = humans.HIGH_INTENSITY
				default:
					intensity = humans.NO_INTENSITY
					fmt.Println("Choose a valid intensity.")
				}
				value := personSelected.Exercise(intensity)
				fmt.Println(value)
			}

		case SLEEP:
			sleep := personSelected.Sleep()
			fmt.Println(sleep)
		case STUDY:
			study := personSelected.Study()
			fmt.Println(study)
		case EAT:
			var foodOption int
			for foodOption == 0 {
				fmt.Println("Which thing should I eat:")
				fmt.Println("1. Dessert")
				fmt.Println("2. Meal")
				fmt.Println("3. Energizer")

				_, err := fmt.Scan(&foodOption)
				if err != nil {
					foodOption = 0
					continue
				}

				var selectedFood food.Food

				switch foodOption {
				case 1:
					selectedFood = food.Dessert{}
				case 2:
					selectedFood = food.Meal{}
				case 3:
					selectedFood = food.Energizer{}
				default:
					foodOption = 0
					fmt.Println("Choose a valid food.")
					continue
				}

				eat := personSelected.Eat(selectedFood)
				fmt.Println(eat)
			}

		case SHOWSTATUS:
			showStatus := personSelected.ShowStatus()
			fmt.Println(showStatus)
		case EXIT:
			fmt.Println("Thanks for play!")
		default:
			fmt.Println("D: action not allowed")
		}
	}
}

// Functionality example
//* person := humans.NewPerson("Wences")
//* child := humans.NewChild("Wences Jr")
// dessert := food.NewDessert("pay","sweet")
// meal := food.NewMeal("meat","tasty")

// fmt.Println(person.Status())
// fmt.Println(person.Exercise(humans.HIGH_INTENSITY))
// fmt.Println(person.Status())
// fmt.Println(person.Eat(meal))
// fmt.Println(person.Eat(dessert))
// fmt.Println(person.Status())

// fmt.Println(child.Status())
// fmt.Println(child.Eat(meal))
// fmt.Println(child.Eat(dessert))
// fmt.Println(child.Status())

package humans

import "trucode.com/interactions/food"

type Humans interface {
	Exercise(level Intensity) string
	Sleep() string
	Study() string
	Eat(food food.Food) string
	ShowStatus() string
}



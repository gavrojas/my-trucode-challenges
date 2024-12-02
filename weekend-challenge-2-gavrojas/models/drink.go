package models

import (
	"fmt"

	datastructures "trucode.app/w2challenge/data_structures"
)

type Drink struct {
	Ingredients datastructures.StringStack
}

func (d *Drink) AddIngredient(ingredient string) {
	d.Ingredients.Push(ingredient)
}

func (d *Drink) RemoveIngredient() (string, error) {
	value, error := d.Ingredients.Pop()

	if error != nil {
		return "", fmt.Errorf("no more ingredients in the list")
	} else {
		return value, nil
	}
}

func (d Drink) ListIngredients() string {
	stringValue, _ := d.Ingredients.ToString()
	return stringValue
}

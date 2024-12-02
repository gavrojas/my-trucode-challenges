package humans

import (
	"fmt"

	"trucode.com/interactions/food"
)

type Child struct {
	*Person
}

func NewChild(name string) *Child {
	return &Child{
		Person: &Person{
			name:       name,
			hunger:     100,
			sleepyness: 10,
			stamina:    80,
		},
	}
}

// metodo eat que cambia para el child
func (c *Child) Eat(food food.Food) string {
	c.hunger += float64(food.GetHungerEffect()) * 0.8
	if c.hunger < 0 {
		c.hunger = 0
	}
	if c.hunger > 100 {
		c.hunger = 100
	}

	c.stamina += float64(food.GetStaminaEffect()) * 1.2
	if c.stamina > 100 {
		c.stamina = 100
	}
	if c.stamina < 0 {
		c.stamina = 0
	}

	c.sleepyness += float64(food.GetSleepinessEffect()) * 1.1
	if c.sleepyness > 100 {
		c.sleepyness = 100
	}
	if c.sleepyness < 0 {
		c.sleepyness = 0
	}

	return fmt.Sprintf("I have eaten. That %s tasted %s!", food.GetName(), food.GetTaste())
}

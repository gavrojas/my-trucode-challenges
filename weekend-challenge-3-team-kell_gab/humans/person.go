package humans

import (
	"fmt"

	"trucode.com/interactions/food"
)

type Person struct {
	name       string
	hunger     float64
	sleepyness float64
	stamina    float64
}

type Intensity string

const (
	LOW_INTENSITY    Intensity = "low"
	MEDIUM_INTENSITY Intensity = "medium"
	HIGH_INTENSITY   Intensity = "high"
	NO_INTENSITY     Intensity = "none"
)

func NewPerson(name string) *Person {
	return &Person{
		name:       name,
		hunger:     80,
		sleepyness: 10,
		stamina:    80,
	}
}

func (p *Person) Exercise(level Intensity) string {
	switch level {
	case LOW_INTENSITY:
		p.stamina -= 10
		p.hunger += 10
	case MEDIUM_INTENSITY:
		p.stamina -= 25
		p.hunger += 30

	case HIGH_INTENSITY:
		p.stamina -= 50
		p.hunger += 60
	}

	if p.hunger > 100 {
		p.hunger = 100
	}
	if p.hunger < 0 {
		p.hunger = 0
	}
	if p.stamina > 100 {
		p.stamina = 100
	}
	if p.stamina < 0 {
		p.stamina = 0
	}

	return fmt.Sprintf("%s has exercised at %s intensity.", p.name, level)

}

func (p *Person) Sleep() string {
	p.hunger += 20
	p.sleepyness = 0
	p.stamina = 100

	if p.hunger > 100 {
		p.hunger = 100
	}
	if p.hunger < 0 {
		p.hunger = 0
	}

	return "I have slept."
}

func (p *Person) Study() string {
	p.hunger += 25
	p.sleepyness += 30
	p.stamina -= 10

	if p.hunger > 100 {
		p.hunger = 100
	}
	if p.hunger < 0 {
		p.hunger = 0
	}
	if p.sleepyness > 100 {
		p.sleepyness = 100
	}
	if p.sleepyness < 0 {
		p.sleepyness = 0
	}
	if p.stamina > 100 {
		p.stamina = 100
	}
	if p.stamina < 0 {
		p.stamina = 0
	}

	return "I have studied"
}

//metodo EAT

func (p *Person) Eat(food food.Food) string {
	p.hunger += food.GetHungerEffect()
	if p.hunger < 0 {
		p.hunger = 0
	}
	if p.hunger > 100 {
		p.hunger = 100
	}

	p.stamina += food.GetStaminaEffect()
	if p.stamina > 100 {
		p.stamina = 100
	}
	if p.stamina < 0 {
		p.stamina = 0
	}

	p.sleepyness += food.GetSleepinessEffect()
	if p.sleepyness > 100 {
		p.sleepyness = 100
	}
	if p.sleepyness < 0 {
		p.sleepyness = 0
	}

	return fmt.Sprintf("I have eaten. That %s tasted %s!", food.GetName(), food.GetTaste())
}

func (p *Person) ShowStatus() string {
	status := fmt.Sprintf("My name is %s, this is my status:\nHunger: %.2f\nSleepyness: %.2f\nStamina: %.2f\n", p.name, p.hunger, p.sleepyness, p.stamina)

	if p.hunger == 100 {
		status += "I am very hungry\n"
	} else if p.hunger == 0 {
		status += "I am satiated\n"
	}

	if p.sleepyness == 100 {
		status += "I need to sleep\n"
	} else if p.sleepyness == 0 {
		status += "I am totally rested\n"
	}

	if p.stamina == 100 {
		status += "I have a lot of energy\n"
	} else if p.stamina == 0 {
		status += "I'm tired\n"
	}

	return status
}

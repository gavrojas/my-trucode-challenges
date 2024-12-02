package food

type Meal struct{}

func (m Meal) GetName() string {
	return "Meal"
}

func (m Meal) GetTaste() string {
	return "savory"
}

func (m Meal) GetHungerEffect() float64 {
	return -50
}

func (m Meal) GetStaminaEffect() float64 {
	return 25
}

func (m Meal) GetSleepinessEffect() float64 {
	return 10
}

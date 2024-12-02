package food

type Dessert struct{}

func (d Dessert) GetName() string {
	return "Dessert"
}

func (d Dessert) GetTaste() string {
	return "sweet"
}

func (d Dessert) GetHungerEffect() float64 {
	return -20
}

func (d Dessert) GetStaminaEffect() float64 {
	return 10
}

func (d Dessert) GetSleepinessEffect() float64 {
	return 0
}

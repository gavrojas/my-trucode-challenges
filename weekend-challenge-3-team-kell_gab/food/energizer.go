package food

type Energizer struct{}

func (e Energizer) GetName() string {
	return "Energizer"
}

func (e Energizer) GetTaste() string {
	return "invigorating"
}

func (e Energizer) GetHungerEffect() float64 {
	return -10
}

func (e Energizer) GetStaminaEffect() float64 {
	return 50
}

func (e Energizer) GetSleepinessEffect() float64 {
	return -50
}

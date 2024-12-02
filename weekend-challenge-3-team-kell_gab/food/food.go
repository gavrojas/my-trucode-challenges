package food

type Food interface {
	GetName() string
	GetTaste() string
	GetHungerEffect() float64
	GetStaminaEffect() float64
	GetSleepinessEffect() float64
}

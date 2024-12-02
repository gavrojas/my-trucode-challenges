package dictionary

type Dictionary struct {
	Word               string  `json:"text"`
	Translation        string  `json:"translation"`
	TotalQuestions     int     `json:"totalQuestion"`
	ContCorrectAnswers int     `json:"correctAnswers"`
	RetencionRate      float64 `json:"retencionRate"`
	RepetitionInterval int     `json:"repetitionInterval"`
}

// UpdateRepetitionInterval ajusta el intervalo de repetición según la tasa de retención
func (d *Dictionary) UpdateRepetitionInterval() {
	switch {
	case d.RetencionRate < 50:
		d.RepetitionInterval = 1 // Repetir pronto (mucha dificultad)
	case d.RetencionRate >= 50 && d.RetencionRate < 80:
		d.RepetitionInterval = 3 // Repetir con moderada frecuencia
	case d.RetencionRate >= 80:
		d.RepetitionInterval = 5 // Repetir menos frecuentemente
	}
}

package users

import (
	"fmt"
	"math"

	"trucode.com/challenge-w5/dictionary"
)

type User struct {
	UserId        int                     `json:"userID"`
	Name          string                  `json:"name"`
	WordsLearning []dictionary.Dictionary `json:"words"`
}

func NewUser(userName string) User {
	return User{
		UserId: 1, //valor user1
		Name:   userName,
		WordsLearning: []dictionary.Dictionary{
			{Word: "learn", Translation: "aprender", TotalQuestions: 0, ContCorrectAnswers: 0, RetencionRate: 0},
			{Word: "run", Translation: "correr", TotalQuestions: 0, ContCorrectAnswers: 0, RetencionRate: 0},
			{Word: "sleep", Translation: "dormir", TotalQuestions: 0, ContCorrectAnswers: 0, RetencionRate: 0},
			{Word: "eat", Translation: "comer", TotalQuestions: 0, ContCorrectAnswers: 0, RetencionRate: 0},
			{Word: "walk", Translation: "caminar", TotalQuestions: 0, ContCorrectAnswers: 0, RetencionRate: 0},
		},
	}
}

func (u *User) OrderWords(correct bool) {
	if correct {
		numMoves := math.Pow(2, float64(u.WordsLearning[0].ContCorrectAnswers))
		if int(numMoves) > len(u.WordsLearning) {
			numMoves = float64(len(u.WordsLearning)) - 1
		}
		for i := 0; i < int(numMoves); i++ {
			u.WordsLearning[i], u.WordsLearning[i+1] = u.WordsLearning[i+1], u.WordsLearning[i]
		}
	} else {
		u.WordsLearning[0], u.WordsLearning[1] = u.WordsLearning[1], u.WordsLearning[0]
	}
}

func (u User) Progress() string {
	result := fmt.Sprintf("Your current progress is:\n")
	result += "Words       | Retention Rate\n"
	for _, value := range u.WordsLearning {
		result += fmt.Sprintf("%-12s| %-5.0f %% \n", value.Word, value.RetencionRate)
	}
	return result
}

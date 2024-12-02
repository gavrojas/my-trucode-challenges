package main

import (
	"fmt"

	"trucode.com/challenge-w5/files"
	"trucode.com/challenge-w5/screen"
	"trucode.com/challenge-w5/users"
)

// Sesión de estudio para un usuario específico
func StudySession(user users.User, filename string, dataUsers *files.InfoUsers) {
	words := user.WordsLearning
	fmt.Printf("What does '%s' mean?\n", words[0].Word)

	var answer string
	fmt.Scan(&answer)

	// Actualizar estadísticas
	words[0].TotalQuestions++
	if answer == words[0].Translation {
		words[0].ContCorrectAnswers++
		fmt.Println("Correct!")
		// Actualizar la tasa de retención
		user.WordsLearning[0].RetencionRate = (float64(words[0].ContCorrectAnswers) / float64(words[0].TotalQuestions)) * 100
		user.OrderWords(true)
	} else {
		fmt.Printf("Incorrect. The correct answer is '%s'.\n", words[0].Translation)
		// Actualizar la tasa de retención
		user.WordsLearning[0].RetencionRate = (float64(words[0].ContCorrectAnswers) / float64(words[0].TotalQuestions)) * 100
		user.OrderWords(false)
	}

	// Guardar los cambios después de cada ciclo
	if err := files.WriteData(filename, *dataUsers); err != nil {
		fmt.Println("Error al guardar los datos:", err)
	}

}

func main() {
	// Nombre del archivo JSON
	filename := "users.json"
	// Leer los datos del archivo JSON
	dataUsers, err := files.ReadData(filename)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}

	fmt.Println("\nThis is Spaced repetition app\n\nIt helps you to learn english \nusing the spaced repetition \nalgorithm!!!\n\nGood luck!")

	var user users.User
	entryOptions := screen.OptionLogin()
	switch entryOptions {
	case screen.REGISTER:
		userName := screen.UserPrompt()
		user = users.NewUser(userName)
		files.AddUserInJson(filename, userName)
		dataUsers, _ = files.ReadData(filename)
		fmt.Println("Welcome to the Spaced Repetition App!")
	case screen.LOGIN:
		user = validateExist(fmt.Errorf("validate"), "", dataUsers)
	}

	selectedOptionLearn := screen.LEARN_UNSELECTED
	for selectedOptionLearn != screen.FINISH {
		selectedOptionLearn = screen.LearnPrompt()
		switch selectedOptionLearn {
		case screen.NEXT_WORD:
			// Iniciar la sesión de estudio para el usuario
			StudySession(user, filename, &dataUsers)
		case screen.PROGRESS:
			dataUsers, _ = files.ReadData(filename)
			fmt.Println(user.Progress())
		}
	}

	fmt.Println("\nThanks for practicing with the \nSpace repetition app \nSee you later!")
}

func validateExist(err error, userName string, dataUsers files.InfoUsers) users.User {
	for err != nil {
		userName = screen.UserLogPrompt()
		_, err = files.FindUser(userName, dataUsers)
		if err != nil {
			fmt.Printf("Invalid user name\n")
		}
	}
	user, _ := files.FindUser(userName, dataUsers)
	fmt.Printf("Hello %s!", user.Name)
	return user
}

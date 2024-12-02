package screen

import "fmt"

// Solicita el nombre del usuario
func UserPrompt() string {
	name := ""
	for name == "" {
		fmt.Print("Please tell me your name?\n")
		if _, err := fmt.Scan(&name); err != nil {
			fmt.Printf("This is not a valid name\n")
		}
	}
	return name
}

// Solicita el nombre del usuario registrado
func UserLogPrompt() string {
	name := ""
	for name == "" {
		fmt.Printf("Please remember me your name?\n")
		if _, err := fmt.Scan(&name); err != nil {
			fmt.Printf("Invalid user name\n")
		}
	}
	return name
}

//Opciones de ingreso
type optionLog int

const (
	LOG_UNSELECTED optionLog = iota
	REGISTER
	LOGIN
)

var optionLoginLabels = map[optionLog]string{
	REGISTER: "Register",
	LOGIN:    "Login",
}

// Indica las opciones de ingreso
func OptionLogin() optionLog {
	return choicePrompt[optionLog]("\nI don't know who you are...\nDid you want to:", optionLoginLabels, LOG_UNSELECTED, REGISTER, LOGIN)
}

type learnOptions int

const (
	LEARN_UNSELECTED learnOptions = iota
	NEXT_WORD
	PROGRESS
	FINISH
)

var learnOptionsLabels = map[learnOptions]string{
	NEXT_WORD: "Go to the next word",
	PROGRESS:  "See my progress",
	FINISH:    "Finish",
}

func LearnPrompt() learnOptions {
	return choicePrompt[learnOptions]("\nWhat did you want to do?", learnOptionsLabels, LEARN_UNSELECTED, NEXT_WORD, PROGRESS, FINISH)
}

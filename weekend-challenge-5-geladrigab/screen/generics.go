package screen

import "fmt"

type choices interface {
	optionLog | learnOptions
}

func choicePrompt[T choices](question string, labelsMap map[T]string, unselected T, options ...T) T {
	userChoice := unselected
	validChoice := false

	for !validChoice {
		fmt.Println(question)
		for _, value := range options {
			fmt.Printf("%d. %s\n", value, labelsMap[value])
		}
		if _, err := fmt.Scan(&userChoice); err != nil || userChoice == unselected || userChoice > options[len(options)-1] {
			fmt.Printf("Invalid choice, please select a valid one\n")
		} else {
			validChoice = true
		}
	}
	return userChoice
}

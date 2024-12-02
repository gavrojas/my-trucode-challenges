package menu

import (
	"fmt"
	"strings"

	datastructures "trucode.app/w2challenge/data_structures"
	models "trucode.app/w2challenge/models"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func CreateClient(c *datastructures.Set) models.Client {
	var name string

	for {
		fmt.Print("\nEnter client name: ")
		fmt.Scanln(&name)
		name = Capitalize(name)

		if c.Exists(name) {
			fmt.Printf("%s's order already exists", name)
			continue
		} else {
			client := models.Client{Name: name}
			fmt.Printf("\n--> Client %s created.\n\n", client.Name)
			return client
		}
	}

}

func CreateDrink() (models.Drink, error) {
	var drink models.Drink
	var option string
	var ingredient string

	for {
		fmt.Println("\nDrink Options:")
		fmt.Println("1. Add Ingredient")
		fmt.Println("2. Remove Last Ingredient")
		fmt.Println("3. Finish Drink and enter order")
		fmt.Println("0. Back to Main Menu")
		fmt.Print("\nChoose an option: ")
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Print("Enter ingredient: ")
			fmt.Scanln(&ingredient)
			drink.AddIngredient(ingredient)
			fmt.Printf("\n--> Added ingredient: %s\nIngredients: ", ingredient)
		case "2":
			_, err := drink.RemoveIngredient()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("\n--> Removed last ingredient.\nIngredients: ")
			}
		case "3":
			if drink.ListIngredients() == "" {
				fmt.Println("No ingredients added to the drink. \nPlease add ingredients:")
				break
				//return models.Drink{}, fmt.Errorf("no ingredients added to the drink")
			}
			fmt.Println("\n-->Drink finished.")
			fmt.Println("Ingredients: ", drink.ListIngredients())
			return drink, nil
		case "0":
			return drink, nil
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func EnterOrder(om *models.OrderManager) {
	client := CreateClient(&om.Clients)
	drink, error := CreateDrink()

	if error != nil {
		fmt.Println(error)
		return
	}

	if drink.ListIngredients() == "" {
		fmt.Println("Error: No drink created. Please create a drink before entering an order.")
		return
	}

	order := models.Order{Client: client, Drink: drink}

	if error := om.Add(&order); error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s entered an order.\n", client.Name)
	}
}

func DeliverOrders(orderManager *models.OrderManager) {
	fmt.Println("Orders: ")
	var counter int
	for {
		order, err := orderManager.Next()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%d. Order for %s is ready to be served with ingredients: %v\n", counter, order.Client.Name, order.Drink.ListIngredients())
		counter++
	}
}

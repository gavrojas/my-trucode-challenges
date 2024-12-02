package main

import (
	"fmt"

	menu "trucode.app/w2challenge/menu"
	models "trucode.app/w2challenge/models"
)

func main() {
	// Create a new orderQueue
	var orderManager models.OrderManager
	orderManager.Initialize()
	var option string

	fmt.Println("\n============================================================================")
	fmt.Println("\t Hello, it's a pleasure to have you here in our coffee shop")

	for {
		fmt.Println("============================================================================")
		fmt.Println("\nYou have the following options:")
		fmt.Println("1. Enter Order")
		fmt.Println("2. Deliver Orders")
		fmt.Println("0. Exit")
		fmt.Print("\nChoose an option: ")
		fmt.Scanln(&option)

		switch option {
		case "1":
			menu.EnterOrder(&orderManager)
		case "2":
			menu.DeliverOrders(&orderManager)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// }
// // Create a client
// client1 := models.Client{Name: "John"}
// fmt.Printf("%s wants to order\n", client1.Name)

// // Create a drink for client1
// var drink1 models.Drink

// // Add ingredients
// drink1.AddIngredient("coffee")
// drink1.AddIngredient("water")
// drink1.AddIngredient("milk")
// fmt.Println(drink1)
// // Remove the last ingredient
// drink1.RemoveIngredient()
// fmt.Println(drink1)
// // Keep adding ingredients
// drink1.AddIngredient("lactose-free milk")
// drink1.AddIngredient("syrup")
// fmt.Println(drink1)

// // Append a new order to the orderQueue
// order1 := models.Order{Client: client1, Drink: drink1}
// fmt.Println(order1)
// orderManager.Add(&order1)
// fmt.Printf("%s placed an order, the required drink has the following ingredients: %v\n", order1.Name, order1.Drink.ListIngredients())

// // Shows error because a client1 order already exists
// fmt.Printf("%s repeats the order, the required drink has the following ingredients: %v\n", order1.Name, order1.Drink.ListIngredients())
// if err := orderManager.Add(&models.Order{Client: client1, Drink: drink1}); err != nil {
// 	fmt.Println(err)
// }

// // Create another client
// client2 := models.Client{Name: "Jane"}

// // Create a drink for client2
// var drink2 models.Drink

// // Add ingredients
// drink2.AddIngredient("coffee")
// drink2.AddIngredient("water")
// drink2.AddIngredient("milk")
// // Remove the last ingredient
// drink2.RemoveIngredient()

// // Append a new order to the orderQueue
// order2 := models.Order{Client: client2, Drink: drink2}
// orderManager.Add(&order2)
// fmt.Printf("%s placed an order, the required drink has the following ingredients: %s\n", order2.Name, order2.Drink.ListIngredients())

// // Deliver the coffees in orderQueue
// nextOrder, _ := orderManager.Next()
// fmt.Printf("Coffee for %s ready to be tasted!!!\n", nextOrder.Name)
// nextOrder, _ = orderManager.Next()
// fmt.Printf("Coffee for %s ready to be tasted!!!\n", nextOrder.Name)
// // Show an error message if there are no more orders to be served
// if _, err := orderManager.Next(); err != nil {
// 	fmt.Println(err)
// }

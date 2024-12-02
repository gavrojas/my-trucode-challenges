package models

import (
	"fmt"

	datastructures "trucode.app/w2challenge/data_structures"
)

type OrderManager struct {
	orders  []Order
	Clients datastructures.Set
}

func (om *OrderManager) Initialize() {
	om.Clients = datastructures.Set{Items: make(map[string]bool)}
}

func (om *OrderManager) Enqueue(order Order) {
	om.orders = append(om.orders, order)
}

func (om *OrderManager) Dequeue() (Order, error) {
	length := len(om.orders)

	if length == 0 {
		return Order{}, fmt.Errorf("no more orders in the stack")
	}

	value := om.orders[0]
	om.orders = om.orders[1:]

	return value, nil
}

func (om *OrderManager) Add(order *Order) error {
	clientName := order.Client.Name
	fmt.Println("Imprimiendo el nombre del cliente: ", clientName)

	if om.Clients.Exists(clientName) {
		return fmt.Errorf("%s's order already exists", clientName)
	}

	if err := om.Clients.Append(clientName); err != nil {
		return err
	}

	om.Enqueue(*order)
	return nil
}

func (om *OrderManager) Next() (Order, error) {
	orderValue, error := om.Dequeue()
	if error != nil {
		return Order{}, fmt.Errorf("no more orders to be served")
	} else {
		clientName := orderValue.Client.Name
		om.Clients.Delete(clientName)
		return orderValue, nil
	}
}

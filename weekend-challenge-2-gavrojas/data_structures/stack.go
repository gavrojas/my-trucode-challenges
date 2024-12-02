package datastructures

import (
	"fmt"
	"strings"
)

type StringStack struct {
	items []string
}

func (ss *StringStack) Push(value string) {
	ss.items = append(ss.items, value)
}

func (ss *StringStack) Pop() (string, error) {
	length := len(ss.items)

	if length == 0 {
		return "", fmt.Errorf("no more items in the stack")
	}

	value := ss.items[length-1]
	ss.items = ss.items[:length-1]
	return value, nil
}

func (ss *StringStack) ToString() (string, error) {
	length := len(ss.items)

	if length == 0 {
		return "", fmt.Errorf("no items in the stack")
	}

	return strings.Join(ss.items, ", "), nil
}

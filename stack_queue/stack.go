package main

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	top        int
	capacity   int
	stackarray []int
}

func push(stack *Stack) error {
	if stack.top == stack.capacity {
		return errors.New("stack capacity limit reached")
	}
	var data int
	fmt.Println("Enter data to enter in stack:")
	fmt.Scan(&data)
	stack.stackarray[stack.top] = data
	stack.top++
	return push(stack)
}

func pop(stack *Stack) (int, error) {
	if stack.top <= 0 {
		return 0, errors.New("No more elements in stack to PopUp")
	} else {
		data := stack.stackarray[stack.top-1]
		stack.top--
		return data, nil
	}
}

func peek(stack *Stack) (int, error) { //peeks shows top element without removing it from stack
	if stack.top <= 0 {
		return 0, errors.New("No elements in stack to peek")
	} else {
		return stack.stackarray[stack.top-1], nil
	}
}

func printStack(stack *Stack) {
	for i := 0; i < stack.top; i++ {
		fmt.Print(stack.stackarray[i], "  ")
	}
	return
}

func main() {
	var stack *Stack
	var capacity int
	var flag string
	fmt.Println("Enter capacity of stack:")
	fmt.Scan(&capacity)
	stack = &Stack{0, capacity, make([]int, capacity)}
	err := push(stack)
	if err != nil {
		fmt.Println(err)
	}
	printStack(stack)
	fmt.Println("Enter 'peek' to see top element & 'pop' to pop element:")
	fmt.Println("pop will display 1st element and delete it from the stack")
	fmt.Scan(&flag)
	flag = strings.ToLower(flag)
	switch flag {
	case "peek":
		value, err := peek(stack)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	case "pop":
		value, err := pop(stack)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	default:
		fmt.Println("You entred wrong option.")
	}
	printStack(stack)
}

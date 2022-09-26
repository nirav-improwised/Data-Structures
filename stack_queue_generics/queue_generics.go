package main

import (
	"errors"
	"fmt"
	"strings"
)

type Queue struct {
	top        int
	end        int
	capacity   int
	queueArray []T
}

type T string //Just Change the data type here to enter Data of different type. Ex: replace string with 'int' to enter integer data.

func push(queue *Queue) error {
	if queue.top == queue.capacity {
		return errors.New("Queue capacity reached")
	}
	var data T
	// var data
	fmt.Println("Enter data:")
	fmt.Scan(&data)
	queue.queueArray[queue.top] = data
	queue.top++
	return push(queue)

}

func pop(queue *Queue) (any, error) {
	if queue.end > (queue.top - 1) {
		return 0, errors.New("No more elements to pop from Queue")
	} else {
		queue.end++
		return queue.queueArray[queue.end-1], nil
	}
}

func peek(queue *Queue) (any, error) {
	if queue.end > (queue.top - 1) {
		return 0, errors.New("No more elements to pop from Queue")
	} else {
		return queue.queueArray[queue.end], nil
	}
}

func printQueue(queue *Queue) {
	for i := queue.end; i < queue.top; i++ {
		fmt.Print(queue.queueArray[i], "  ")
	}
	return
}

func main() {
	var queue *Queue
	var capacity int
	var flag string
	fmt.Println("Enter capacity of Queue:")
	fmt.Scan(&capacity)
	queue = &Queue{0, 0, capacity, make([]T, capacity)}
	err := push(queue)
	if err != nil {
		fmt.Println(err)
	}
	printQueue(queue)
	fmt.Println("\nEnter 'peek' or 'pop' to see elements entered in FIFO order:")
	fmt.Println("Pop will display and delete the elements in FIFO order:")
	fmt.Scan(&flag)
	flag = strings.ToLower(flag)
	switch flag {
	case "peek":
		value, err := peek(queue)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	case "pop":
		value, err := pop(queue)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	default:
		fmt.Println("You entered inappropriate choice")
	}
	printQueue(queue)

}

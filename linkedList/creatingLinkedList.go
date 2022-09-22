package main

import (
	"fmt"
)

func create(n *node) {
	var number int
	fmt.Println("Enter a number to add to linked list or -999 to terminate list:")
	fmt.Scan(&number)
	if number == -999 {
		n = nil
		return
	} else {
		n.snum = number
		n.next = &node{}
		create(n.next)
	}
}
func printList(n *node) {
	if n.next == nil {
		return
	} else {
		fmt.Print(n.snum, "\t")
		printList(n.next)
	}
}

type node struct {
	snum int
	next *node
}

func main() {
	var head *node
	head = &node{}
	create(head)
	printList(head)
}

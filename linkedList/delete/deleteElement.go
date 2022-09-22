package main

import "fmt"

type node struct {
	snum int
	next *node
}

func create(n *node) {
	var number int
	fmt.Println("Enter a number to add to list or -999 to terminate list:")
	fmt.Scan(&number)
	if number == -999 {
		n = nil
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

func deleteElement(n *node, key int) {
	if n.snum == key { //if first node is to be deleted
		*n = *n.next
		return
	}
	previousNode := findKeyNode(n, key)
	if previousNode.next.next != nil { //for deleting nodes in between first and last
		previousNode.next = previousNode.next.next
	} else { //for deleting last node
		previousNode.next = nil
	}
}

func findKeyNode(n *node, key int) *node {
	if n.next.snum == key {
		return n
	} else {
		return findKeyNode(n.next, key)
	}
}

func main() {
	var key int
	var head *node
	head = &node{}
	create(head)
	printList(head)
	fmt.Println("\nEnter a number to delete:")
	fmt.Scan(&key)
	deleteElement(head, key)
	// head = head.next
	printList(head)

}

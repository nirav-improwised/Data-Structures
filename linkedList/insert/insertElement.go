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
		// n = nil
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

func length(n *node) int {
	if n.next == nil {
		return 0
	} else {
		return 1 + length(n.next)
	}
}

func insertElement(n *node) {
	var pos int
	var key int
	fmt.Println("\nEnter a number to insert:")
	fmt.Scan(&key)
	fmt.Println("Enter position to enter number at:")
	fmt.Scan(&pos)
	if pos < 1 || pos > ((length(n))+1) {
		fmt.Println("Enter a valid index please: ")
		insertElement(n)
	}
	if pos == 1 { //if number is to be inserted at first element
		var newNode *node
		newNode = &node{key, n}
		n = newNode
		fmt.Println("newNode add:", newNode)
		fmt.Println("newNode.next add(hsoube be =to old head)", newNode.next)
		fmt.Println("new head add(=to newNode)", n)
		printList(n)
		return
	}
	previousNode := findKeyNode(n, pos)
	// if previousNode.next == nil { //to add an element at last position
	// 	newNode := &node{key, nil}
	// 	fmt.Println(":::", previousNode.next.snum)
	// 	fmt.Println(previousNode.next.next)
	// 	previousNode.next.next = newNode
	// 	fmt.Println(previousNode.next.next)
	// } else {
	newNode := &node{key, previousNode.next}
	previousNode.next = newNode
	// }
	printList(n)

}

var loopCtrl = 1

func findKeyNode(n *node, pos int) *node {
	var keynode *node
	keynode = n
	if loopCtrl < pos-1 {
		loopCtrl++
		keynode = findKeyNode(n.next, pos)
	}
	return keynode
}

func main() {
	var head *node
	head = &node{}
	create(head)
	printList(head)
	fmt.Println("\nOld head add:", head)
	insertElement(head)
	fmt.Println("New head add:", head)
}

package main

import "fmt"

// Cell represents node inside a LinkedList
type Cell struct {
	data string
	next *Cell
}

// LinkedList represents a LinkedList starting with a sentinel
type LinkedList struct {
	sentinel *Cell
}

// makeLinkedList creates an empty LinkedList setting up the sentinel node
func makeLinkedList() LinkedList {
	return LinkedList{sentinel: &Cell{"Sentinel", nil}}
}

// addAfter adds the cell `after` after the current cell.
func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	me.next = after
}

// deleteAfter deletes the cell after the current cell.
func (me *Cell) deleteAfter() *Cell {
	if me.next == nil {
		return &Cell{data: "", next: nil}
	}
	toDelete := me.next
	me.next = toDelete.next
	return toDelete
}

// addRange appends at the end of the linked list a range of cells with value
// specified by `values`.
func (list *LinkedList) addRange(values []string) {
	cell := list.sentinel
	for cell.next != nil {
		cell = cell.next
	}
	for _, v := range values {
		newCell := Cell{data: v, next: nil}
		cell.addAfter(&newCell)
		cell = cell.next
	}
}

// toString returns a string with all the values from the list separated by
// `separator`.
func (list *LinkedList) toString(separator string) string {
	cell := list.sentinel
	var s string
	for cell != nil {
		s += cell.data
		s += separator
		cell = cell.next
	}
	return s
}

// length returns the length of a linked list.
func (list *LinkedList) length() int {
	cellCount := 0
	cell := list.sentinel
	for cell != nil {
		cellCount++
		cell = cell.next
	}
	return cellCount
}

// isEmpty checks if a list is empty or not.
func (list *LinkedList) isEmpty() bool {
	return list.sentinel.next == nil
}

// push creates a Cell holding the data `value` and adds the cell after the sentinel.
func (list *LinkedList) push(value string) {
	list.sentinel.addAfter(&Cell{data: value})
}

// pop pops the item on top of the stack
func (list *LinkedList) pop() string {
	deletedCell := list.sentinel.deleteAfter()
	return deletedCell.data
}

func main() {
	// Make a list from a slice of values.
	greek_letters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := makeLinkedList()
	list.addRange(greek_letters)
	fmt.Println(list.toString(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := makeLinkedList()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.isEmpty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.toString(" "))
	}
}

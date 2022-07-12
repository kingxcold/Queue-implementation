package main

import "fmt"

func main() {
	var myQueue_slice Queue
	var myQueue_ll QueueLL

	myQueue_slice.Enqueue(5)
	myQueue_slice.Enqueue(3)

	myQueue_ll.Enqueue(5)
	myQueue_ll.Enqueue(3)

	fmt.Println("************* Slice *************")
	myQueue_slice.Display()
	fmt.Println("***************************************")

	fmt.Println()

	fmt.Println("************* Linked list *************")
	myQueue_ll.Display()
	fmt.Println("***************************************")
}

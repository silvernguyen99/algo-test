package main

import (
	"fmt"
	question "sky_mavis/question"
)

func main() {

	///////////////// question 2 /////////////////
	dl := &question.DoubleLinkedList{}
	dl.PushBack(8)

	dlStr := dl.String()
	fmt.Print("list after PushBack(8): " + dlStr + "\n")

	dl.PushBack(15)
	dlStr = dl.String()
	fmt.Print("list after PushBack(15): " + dlStr + "\n")

	dl.PushFront(20)
	dlStr = dl.String()
	fmt.Print("list after PushFront(20): " + dlStr + "\n")

	dl.PushBack(30)
	dlStr = dl.String()
	fmt.Print("list after PushBack(30): " + dlStr + "\n")

	// PopBack
	v1, err := dl.PopBack()
	if err != nil {
		fmt.Print("error PopBack: ", err)
	}

	dlStr = dl.String()
	fmt.Printf("list after PopBack(): "+dlStr+" -- value pop: %d"+"\n", v1)

	v2, err := dl.PopFront()
	if err != nil {
		fmt.Print("error PopFront: ", err)
	}

	dlStr = dl.String()
	fmt.Printf("list after PopFront(): "+dlStr+" -- value pop: %d"+"\n", v2)

	dl.PopBack()
	dl.PopFront()

	// expectation: print error "list is empty"
	_, err = dl.PopFront()
	if err != nil {
		fmt.Print("error PopFront: ", err)
	}

	// expectation: print error "list is empty"
	_, err = dl.PopBack()
	if err != nil {
		fmt.Print("\nerror PopBack: ", err)
	}
}

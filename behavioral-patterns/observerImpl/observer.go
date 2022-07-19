package main

import "fmt"

//Define the interface for an observer type
type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

//ToDo: To conform to the interface, it must have an opUpdate function
func (dl *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", dl.Name, "got data change:", data)
}

package main

import "fmt"

func main() {

	//This is a PULL model
	//Here the consumer has control over how and when the data is accessed
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}

}

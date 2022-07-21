package main

import "fmt"

func main() {
	//Note that by using a callback, we have created an isolation between the code that iterates over all the books,
	//and the underlying implementation of the storage of the book objects in the library
	// The callback function has no control over how many times it called. It is therefore a PUSH orientation

	//Using IterateBooks via a callback function
	lib.IterateBooks(myBookCallback)

	//Using IterateBooks via an anonymous function
	lib.IterateBooks(func(b Book) error {
		fmt.Println("Book author:", b.Author)
		return nil
	})

}

//This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}

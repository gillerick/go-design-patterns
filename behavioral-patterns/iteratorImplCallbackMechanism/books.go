package main

import "fmt"

type BookType int

//Predefined book types
const (
	HardCopy BookType = iota
	SoftCover
	PaperBack
	EBook
)

//Book represents data about a book
type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

//Library holds the collection of books
type Library struct {
	Collection []Book
}

//IterateBooks calls on the given callback function for each book in the collection
func (l *Library) IterateBooks(f func(Book) error) {
	var err error
	for _, b := range l.Collection{
		err = f(b)
		if err != nil{
			fmt.Println("Error encountered:", err)
		}
	}
}

//Library structure to hold a set of Books
var lib *Library = &Library{Collection: []Book{
	{
		Name:      "War and Peace",
		Author:    "Leo Tolstoy",
		PageCount: 864,
		Type:      HardCopy,
	},
	{
		Name:      "Crime and Punishment",
		Author:    "Leo Tolstoy",
		PageCount: 1225,
		Type:      SoftCover,
	},
	{
		Name:      "Brave New World",
		Author:    "Aldous Huxley",
		PageCount: 325,
		Type:      PaperBack,
	},
	{
		Name:      "Catcher in the Rye",
		Author:    "J.D. Salinger",
		PageCount: 206,
		Type:      HardCopy,
	},
	{
		Name:      "To Kill a Mockingbird",
		Author:    "Harper Lee",
		PageCount: 399,
		Type:      PaperBack,
	},
	{
		Name:      "The Grapes of Wrath",
		Author:    "John Steinbeck",
		PageCount: 464,
		Type:      HardCopy,
	},
	{
		Name:      "Withering Heights",
		Author:    "Emily Bronte",
		PageCount: 288,
		Type:      EBook,
	},
}}

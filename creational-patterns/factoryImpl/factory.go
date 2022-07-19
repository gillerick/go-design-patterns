package main

import "fmt"

//A factory function that creates the specified publication
func newPublication(pubType string, name string, pages int, pub string) (iPublication, error) {
	if pubType == "newspaper" {
		return createNewspaper(name, pages, pub), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pages, pub), nil
	}
	return nil, fmt.Errorf("No such publication type.")
}

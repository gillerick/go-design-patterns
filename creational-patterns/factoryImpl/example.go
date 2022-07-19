package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Times", 50, "The Times")
	mag2, _ := newPublication("newspaper", "Daily Nation", 62, "Nation Media Group")
	pubDetails(mag1)
	pubDetails(mag2)
}

/*The code that creates each publication doesn't have to know the specific class type for each publication
We can now add more types of objects and as long as they implement the iPublication interface, we can always
change how the construction process of each object works without effecting the code where the objects are created and used
*/
func pubDetails(pub iPublication) {
	fmt.Printf("_______________________\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("____________________________________\n")
}

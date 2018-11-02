package main

import "fmt"

func main() {
	fmt.Print("My weigth on the surface of march is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.\n")

	// secopnd way to write

	fmt.Println("My weigth on the surface of march is", 149.0 * 0.3783, "lbs, and I would be", 41 * 365 / 687, "years old.")

	// third way to write it
	fmt.Printf("My weigth on the surface of march is %v lbs,", 149.0 * 0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41 * 365 / 687)

	// formating multiple variables
	fmt.Printf("My weigth on the surface of %v lbs is %v.\n", "earth", 149.0)

	// formating for align text with space caracters
	// negative value complete with space at rigth and positive value complet with space at left
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94 )
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	// it's quick difficult to padding with another character than space
	
}
package main

import "fmt"

// If you want to test your code, you may test it here...
func main() {
	// Task 1.a
	persons := []Person{
		{Name: "Bob", CriminalScore: 10},
		{Name: "Rian", CriminalScore: 20},
		{Name: "Liam", CriminalScore: 30},
		{Name: "Alice", CriminalScore: 30},
		{Name: "Tomi", CriminalScore: 40},
		{Name: "Xavy", CriminalScore: 60},
		{Name: "Sero", CriminalScore: 70},
		{Name: "Poni", CriminalScore: 80},
		{Name: "Dani", CriminalScore: 90},
		{Name: "Charlie", CriminalScore: 100},
	}

	persons2 := []Person{
		{Name: "Bob", CriminalScore: 10},
		{Name: "Rian", CriminalScore: 20},
		{Name: "Liam", CriminalScore: 30},
		{Name: "Alice", CriminalScore: 30},
	}

	persons3 := []Person{}
	onTransport, waiting := LastDayInJail(persons, "Poni")
	onTransport2, waiting2 := LastDayInJail(persons2, "")
	onTransport3, waiting3 := LastDayInJail(persons3, "")

	fmt.Println(onTransport, waiting)
	fmt.Println(onTransport2, waiting2)
	fmt.Println(onTransport3, waiting3)

	// Task 2.a
	// _ := RotateImage(nil)

	// Task 2.b
	// RunRotateActualImage()

	// Task 3.a
	// _ := RobotTranslatorV2("")
}

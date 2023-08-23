package main

import "fmt"

// If you want to test your code, you may test it here...
func main() {
	// Task 1.a
	persons := []Person{
		{Name: "Bob", CriminalScore: 10},
		{Name: "Rian", CriminalScore: 30},
		{Name: "Liam", CriminalScore: 20},
		{Name: "Alice", CriminalScore: 15},
		{Name: "Tomi", CriminalScore: 40},
		{Name: "Sero", CriminalScore: 35},
		{Name: "Xavy", CriminalScore: 25},
		{Name: "Poni", CriminalScore: 80},
		{Name: "Dani", CriminalScore: 60},
		{Name: "Charlie", CriminalScore: 55},
	}
	onTransport, waiting := LastDayInJail(persons, "")
	fmt.Println(onTransport, waiting)

	// Task 2.a
	// _ := RotateImage(nil)

	// Task 2.b
	// RunRotateActualImage()

	// Task 3.a
	// _ := RobotTranslatorV2("")
}

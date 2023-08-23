package main

import "strings"

type Person struct {
	Name          string
	CriminalScore int
}

// Task 1.a
func LastDayInJail(criminals []Person, chosenPerson string) (onTransport []Person, waiting []Person) {
	// Write your code here
	bubbleSort(criminals)
	// --------------------
	return nil, nil
	// --------------------
}

func searchChosenPerson(criminals []Person, chosenPerson string) int {
	for i := 0; i < len(criminals); i++ {
		if strings.ToLower(criminals[i].Name) == strings.ToLower(chosenPerson) {
			return i
		}
	}
	return -1
}

func bubbleSort(criminals []Person) {
	swapped := false

	for {
		for i := 0; i < len(criminals)-1; i++ {
			if criminals[i].CriminalScore > criminals[i+1].CriminalScore || criminals[i].Name > criminals[i+1].Name {
				criminals[i], criminals[i+1] = criminals[i+1], criminals[i]
				swapped = true
			} else {
				swapped = false
			}
		}
		if swapped == false {
			break
		}
	}

}

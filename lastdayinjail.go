package main

import "strings"

type Person struct {
	Name          string
	CriminalScore int
}

// Task 1.a
func LastDayInJail(criminals []Person, chosenPerson string) (onTransport []Person, waiting []Person) {
	// Write your code here
	released := []Person{}
	lengthCriminals := len(criminals)

	if lengthCriminals == 0 {
		return []Person{}, []Person{}
	}

	sortCriminals(criminals, lengthCriminals)

	setReleasedCriminals(lengthCriminals, &released, criminals, chosenPerson)

	releaseProcess(released, &onTransport, &waiting)

	// --------------------
	return onTransport, waiting
	// --------------------
}

func setReleasedCriminals(lengthCriminals int, released *[]Person, criminals []Person, chosenPerson string) {
	defaultReleasedCount := 5
	if lengthCriminals >= defaultReleasedCount {
		*released = append(*released, criminals[:defaultReleasedCount]...)
	} else {
		*released = append(*released, criminals...)
	}

	chosenPersonIndex := searchChosenPerson(criminals, chosenPerson)
	if (chosenPerson != "") && (chosenPersonIndex != -1) {
		*released = append(*released, criminals[chosenPersonIndex])
	}

}

func releaseProcess(released []Person, onTransport *[]Person, waiting *[]Person) {
	lengthReleased := len(released)
	defaultVehicleSeats := 3
	if lengthReleased >= defaultVehicleSeats {
		*onTransport = append(*onTransport, released[:defaultVehicleSeats]...)
		*waiting = append(*waiting, released[defaultVehicleSeats:]...)
	} else {
		*onTransport = append(*onTransport, released...)
		*waiting = append(*waiting, []Person{}...)
	}
}

func searchChosenPerson(criminals []Person, chosenPerson string) int {
	for i := 0; i < len(criminals); i++ {
		if strings.ToLower(criminals[i].Name) == strings.ToLower(chosenPerson) {
			return i
		}
	}
	return -1
}

func sortCriminals(criminals []Person, lengthCriminals int) {
	swapped := false

	for {
		for i := 0; i < lengthCriminals-1; i++ {
			if criminals[i].CriminalScore > criminals[i+1].CriminalScore {
				temp := criminals[i]
				criminals[i] = criminals[i+1]
				criminals[i+1] = temp
				swapped = true
			} else if criminals[i].CriminalScore == criminals[i+1].CriminalScore && criminals[i].Name > criminals[i+1].Name {
				temp := criminals[i]
				criminals[i] = criminals[i+1]
				criminals[i+1] = temp
				swapped = false
			} else {
				swapped = false
			}
		}
		if swapped == false {
			break
		}
	}

}

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

	sort(criminals)

	if len(criminals) >= 5 {
		released = append(released, criminals[:5]...)
	} else {
		released = append(released, criminals...)
	}

	chosenPersonIndex := searchChosenPerson(criminals, chosenPerson)
	if (chosenPerson != "") && (chosenPersonIndex != -1) {
		released = append(released, criminals[chosenPersonIndex])
	}

	if len(released) >= 3 {
		onTransport = append(onTransport, criminals[:3]...)
		waiting = append(waiting, released[3:]...)
	} else {
		onTransport = append(onTransport, criminals...)
		waiting = append(waiting, []Person{}...)
	}

	// --------------------
	return onTransport, waiting
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

func sort(criminals []Person) {
	swapped := false

	for {
		for i := 0; i < len(criminals)-1; i++ {
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

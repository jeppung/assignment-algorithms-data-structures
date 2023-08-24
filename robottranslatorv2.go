package main

import (
	"fmt"
	"strings"
)

// Task 3.a
func RobotTranslatorV2(cmd string) string {
	// Write your code here
	temp := []string{}
	filteredCmd := []string{}
	var flag int = 0
	var count int

	for i := flag; i < len(cmd); i++ {
		letterCase := strings.ToUpper(string(cmd[i]))
		if letterCase == "X" {
			if len(filteredCmd)-1 != -1 {
				filteredCmd = filteredCmd[:len(filteredCmd)-1]
				flag = 0
			}
		} else {
			filteredCmd = append(filteredCmd, letterCase)
		}
	}

	for i := 0; i < len(filteredCmd); i++ {
		switch string(filteredCmd[i]) {
		case "R":
			if i != 0 && string(filteredCmd[i-1]) == "R" {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move right %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move right %d time", count))
			}
		case "L":
			if i != 0 && string(filteredCmd[i-1]) == "L" {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move left %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move left %d time", count))
			}
		case "A":
			if i != 0 && string(filteredCmd[i-1]) == "A" {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move advance %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move advance %d time", count))
			}
		default:
			return "Invalid command"
		}
	}

	// --------------------
	return strings.Join(temp, "\n")
	// --------------------
}

package main

import (
	"fmt"
	"strings"
)

// Task 3.a
func RobotTranslatorV2(cmd string) string {
	// Write your code here
	temp := []string{}
	temp2 := []string{}
	temp3 := []string{}
	var flag int = 0
	var count int

	temp2 = strings.Split(cmd, "")
	for i := flag; i < len(cmd); i++ {
		if temp2[i] == "X" {
			fmt.Println(temp3)
			if len(temp3)-1 != -1 {
				temp3 = temp3[:len(temp3)-1]
				flag = 0
			}
		} else {
			temp3 = append(temp3, temp2[i])
		}
	}

	for i := 0; i < len(temp3); i++ {
		switch string(temp3[i]) {
		case "R":
			if i != 0 && string(temp3[i-1]) == "R" {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move right %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move right %d times", count))
			}
		case "L":
			if i != 0 && string(temp3[i-1]) == "L" {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move left %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move left %d times", count))
			}
		case "A":
			if i != 0 && string(temp3[i-1]) == "A" {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move advance %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move advance %d times", count))
			}
		}
	}

	fmt.Println(temp)
	// --------------------
	return ""
	// --------------------
}

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
	rightCmd := "R"
	leftCmd := "L"
	advanceCmd := "A"
	var count int

	filterCmd(cmd, &filteredCmd)

	for i := 0; i < len(filteredCmd); i++ {
		switch string(filteredCmd[i]) {
		case rightCmd:
			processCommand(&filteredCmd, rightCmd, "right", &count, &temp, &i)
		case leftCmd:
			processCommand(&filteredCmd, leftCmd, "left", &count, &temp, &i)
		case advanceCmd:
			processCommand(&filteredCmd, advanceCmd, "advance", &count, &temp, &i)
		default:
			return "Invalid command"
		}
	}

	// --------------------
	return strings.Join(temp, "\n")
	// --------------------
}

func processCommand(filteredCmd *[]string, cmd string, cmdName string, count *int, temp *[]string, currIndex *int) {
	if *currIndex != 0 && string((*filteredCmd)[*currIndex-1]) == cmd {
		*count++
		(*temp)[len(*temp)-1] = fmt.Sprintf("Move %s %d times", cmdName, *count)
	} else {
		*count = 1
		*temp = append(*temp, fmt.Sprintf("Move %s %d time", cmdName, *count))
	}
}

func filterCmd(cmd string, filteredCmd *[]string) {
	flag := 0
	cancelCmd := "X"
	for i := flag; i < len(cmd); i++ {
		letterCase := strings.ToUpper(string(cmd[i]))
		if letterCase == cancelCmd {
			if len(*filteredCmd)-1 != -1 {
				*filteredCmd = (*filteredCmd)[:len(*filteredCmd)-1]
				flag = 0
			}
		} else {
			*filteredCmd = append(*filteredCmd, letterCase)
		}
	}
}

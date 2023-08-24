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
			if i != 0 && string(filteredCmd[i-1]) == rightCmd {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move right %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move right %d time", count))
			}
		case leftCmd:
			if i != 0 && string(filteredCmd[i-1]) == leftCmd {
				count++
				temp[len(temp)-1] = fmt.Sprintf("Move left %d times", count)
			} else {
				count = 1
				temp = append(temp, fmt.Sprintf("Move left %d time", count))
			}
		case advanceCmd:
			if i != 0 && string(filteredCmd[i-1]) == advanceCmd {
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

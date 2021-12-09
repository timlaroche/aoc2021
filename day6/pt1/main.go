package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	state := []int{}
	s.Scan()
	splitInput := regexp.MustCompile(",").Split(s.Text(), -1)

	for _, val := range splitInput {
		intVal, _ := strconv.Atoi(val)
		state = append(state, intVal)
	}

	for i := 0; i < 80; i++ {
		for i, val := range state {
			newVal := val - 1
			if newVal < 0 {
				state[i] = 6
				state = append(state, 8)
			} else {
				state[i] = newVal
			}
		}
	}

	fmt.Println(len(state))
}

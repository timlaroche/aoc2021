package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	// assume 12 bit strings
	valueOne := [12]int{}
	valueZero := [12]int{}
	gamma := ""   // most common
	epsilon := "" // least common

	for s.Scan() {
		text := s.Text()
		for i := 0; i < 12; i++ {
			value, _ := strconv.Atoi(string(text[i]))
			if value == 1 {
				valueOne[i] += 1
			} else {
				valueZero[i] += 1
			}
		}
	}

	for i := 0; i < 12; i++ {
		if valueOne[i] > valueZero[i] {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	for i := 0; i < 12; i++ {
		if valueOne[i] > valueZero[i] {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}

	fmt.Println(gamma)
	fmt.Println(epsilon)
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	// fmt.Println(gammaDecimal)
	// fmt.Println(epsilonDecimal)
	fmt.Print(gammaDecimal * epsilonDecimal)
}

package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	max := -99999
	costs := [10000]int{}

	s.Scan()

	for _, val := range regexp.MustCompile(",").Split(s.Text(), -1) {
		intVal, _ := strconv.Atoi(val)
		if intVal > max {
			max = intVal
		}
	}

	f, _ = os.Open("../input")
	s = bufio.NewScanner(f)

	for i := 0; i < max; i++ {
		for _, val := range regexp.MustCompile(",").Split(s.Text(), -1) {
			intVal, _ := strconv.Atoi(val)
			costs[i] += intVal - i
		}
	}
}

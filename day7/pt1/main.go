package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	max := -99999
	costs := [10000]float64{}
	hasCrab := map[int]bool{}

	s.Scan()

	for _, val := range regexp.MustCompile(",").Split(s.Text(), -1) {
		intVal, _ := strconv.Atoi(val)
		hasCrab[intVal] = true
		if intVal > max {
			max = intVal
		}
	}

	fmt.Println(max)

	f, _ = os.Open("../input")
	s = bufio.NewScanner(f)
	s.Scan()

	for i := 0; i < max; i++ {
		if hasCrab[i] {
			for _, val := range regexp.MustCompile(",").Split(s.Text(), -1) {
				intVal, _ := strconv.Atoi(val)
				costs[i] += math.Abs(float64(intVal - i))
			}
		}
	}

	min := 999999999.0
	for i := 0; i < max; i++ {
		if hasCrab[i] {
			if costs[i] < min {
				min = costs[i]
			}
		}
	}
	fmt.Println(min)
}

// 1698 too low

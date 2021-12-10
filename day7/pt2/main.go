package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

// is there a prolbme here?
func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	max := -99999
	costs := [10000]int{}
	hasCrab := map[int]bool{}

	s.Scan()

	for _, val := range regexp.MustCompile(",").Split(s.Text(), -1) {
		intVal, _ := strconv.Atoi(val)
		hasCrab[intVal] = true
		if intVal > max {
			max = intVal
		}
	}

	f, _ = os.Open("../input")
	s = bufio.NewScanner(f)
	s.Scan()

	for i := 0; i < max; i++ {

		for _, val := range regexp.MustCompile(",").Split(s.Text(), -1) {
			intVal, _ := strconv.Atoi(val)
			costs[i] += cost(int(math.Abs(float64(intVal - i))))
		}
	}

	fmt.Println(costs)

	min := 999999999
	for i := 0; i < max; i++ {
		if costs[i] < min {
			min = costs[i]
		}

	}
	fmt.Println(min)
}

func cost(difference int) int {
	cost := 0
	for i := 1; i < difference+1; i++ {
		cost += i
	}
	return cost
}

// 1698 too low
// 85015849 too high

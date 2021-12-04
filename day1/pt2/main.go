package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input")
	s := bufio.NewScanner(f)

	count := 0
	prevWindow := [3]int{0, 0, 0}
	// First window
	s.Scan()
	for i := 0; i < 3; i++ {
		value, _ := strconv.Atoi(s.Text())
		prevWindow[i] = value
		s.Scan()
	}

	for s.Scan() {
		value, _ := strconv.Atoi(s.Text())
		//bump
		newWindow := [3]int{}
		newWindow[0] = prevWindow[1]
		newWindow[1] = prevWindow[2]
		newWindow[2] = value

		// comparision
		prevWindowSum := sumArray(prevWindow[:]...)
		newWindowSum := sumArray(newWindow[:]...)

		if newWindowSum > prevWindowSum {
			count++
		}
		prevWindow = newWindow
	}
	fmt.Println(count)
}

func sumArray(nums ...int) int {
	//use slice to sum array
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

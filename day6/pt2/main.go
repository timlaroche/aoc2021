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

	fish := [9]int{}
	s.Scan()
	splitInput := regexp.MustCompile(",").Split(s.Text(), -1)

	for _, val := range splitInput {
		intVal, _ := strconv.Atoi(val)
		fish[intVal] += 1
	}

	//fmt.Println(fish)

	noOfDay := 256

	// n0
	// n1
	// n2
	// n3
	// n4
	// n5
	// n6
	// n7

	for i := 0; i < noOfDay; i++ {
		fmt.Println(fish)
		oldFishCopy := fish
		fish[0] = oldFishCopy[1]
		fish[1] = oldFishCopy[2]
		fish[2] = oldFishCopy[3]
		fish[3] = oldFishCopy[4]
		fish[4] = oldFishCopy[5]
		fish[5] = oldFishCopy[6]
		fish[6] = oldFishCopy[7] + oldFishCopy[0]
		fish[7] = oldFishCopy[8]
		fish[8] = oldFishCopy[0]
	}

	sum := 0
	for _, i := range fish {
		sum += i
	}
	fmt.Println(sum)
}

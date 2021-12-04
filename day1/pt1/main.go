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

	//First Value
	s.Scan()
	prevValue, _ := strconv.Atoi(s.Text())

	for s.Scan() {
		value, _ := strconv.Atoi(s.Text())
		if value > prevValue {
			count++
		}
		prevValue = value
	}

	fmt.Println(count)
}

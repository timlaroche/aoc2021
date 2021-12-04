package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, _ := os.Open("../input")
	// s := bufio.NewScanner(f)

	// assume 12 bit strings
	value := [12]byte{}

	// 13 so we can index using i
	for i := 0; i < 12; i++ {
		oneCommon := doOnce(i)
		if oneCommon {
			value[i] = 1
		} else {
			value[i] = 0
		}
	}
	fmt.Println(value)
}

func doOnce(index int) bool {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	one_count := 0
	zero_count := 0

	for s.Scan() {
		// assume every string is 12 bits long
		text := s.Text()
		fmt.Println(text)
		if text[index] == 1 {
			one_count += 1
		} else {
			zero_count += 0
		}
	}
	if one_count > zero_count {
		fmt.Println(one_count)
		return true
	}
	return false
}

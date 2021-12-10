package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("../input copy")
	s := bufio.NewScanner(f)

	sum := 0
	for s.Scan() {
		digitCode := regexp.MustCompile("\\|").Split(s.Text(), -1)
		digits := regexp.MustCompile(" ").Split(digitCode[1], -1)
		digits = removeEmpty(digits)
		fmt.Println(digits)

		var sb string
		for _, digit := range digits {
			val := countChars(digit)
			signals := map[rune]bool{}
			for _, signal := range digit {
				signals[signal] = true
			}
			switch val {
			case 2: // 2 segments on = 1
				sb += "1"
			case 3: // 3 segments on = 7
				sb += "7"
			case 4: // 4 segments on = 4
				sb += "4"
			case 7: // 7 segments on = 8
				sb += "8"
			case 5: // 5 segemnts on = 2 or 3 or 5
				for _, signal := range digit {
					signals[signal] = true
				}
				if signals['e'] && signals['b'] {
					sb += "5"
				}
				if signals['g'] && signals['a'] {
					sb += "2"
				}
				if signals['b'] && signals['a'] {
					sb += "3"
				}
			case 6:
				for _, signal := range digit {
					signals[signal] = true
				}
				if signals['e'] && signals['f'] && signals['a'] {
					sb += "9"
				}
				if signals['f'] && signals['g'] && signals['e'] {
					sb += "6"
				}
				if signals['a'] && signals['g'] && signals['e'] {
					sb += "0"
				}
			}
		}
		intVal, _ := strconv.Atoi(sb)
		fmt.Println(intVal)
		sum += intVal
	}
	fmt.Println(sum)
}

func countChars(s string) int {
	count := 0
	for range s {
		count++

	}
	return count
}

func removeEmpty(s []string) []string {
	var r []string
	for _, val := range s {
		if val != "" {
			r = append(r, val)
		}
	}
	return r
}

// 1413 too low
// 5398 too low
// 16486 too low
// 57650845571

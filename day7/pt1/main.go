package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	ones := 0
	fours := 0
	sevens := 0
	eights := 0

	for s.Scan() {
		digitCode := regexp.MustCompile("\\|").Split(s.Text(), -1)
		fmt.Println(digitCode)
		digits := regexp.MustCompile(" ").Split(digitCode[1], -1)
		digits = removeEmpty(digits)

		for _, digit := range digits {
			val := countChars(digit)
			switch val {
			case 2: // 2 segments on = 1
				ones++
			case 3: // 3 segments on = 7
				sevens++
			case 4: // 4 segments on = 4
				fours++
			case 7: // 7 segments on = 8
				eights++
			}
		}
	}

	fmt.Println(ones + sevens + fours + eights)
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

// 1698 too low

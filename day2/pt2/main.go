package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("./input")
	s := bufio.NewScanner(f)

	horizontal := 0
	depth := 0
	aim := 0

	for s.Scan() {
		text := s.Text()
		regex := regexp.MustCompile(" ")

		split := regex.Split(text, -1)
		value, _ := strconv.Atoi(split[1])

		if split[0] == "forward" {
			horizontal += value
			if aim > 0 {
				depth += value * aim
			}
		}
		if split[0] == "down" {
			aim += value
		}
		if split[0] == "up" {
			aim -= value
		}
	}
	fmt.Println(horizontal * depth)
}

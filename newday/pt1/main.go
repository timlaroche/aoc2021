package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []rune

func (s *stack) push(v rune) {
	*s = append(*s, v)
}

func (s *stack) pop() rune {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *stack) peek() rune {
	value := (*s)[len(*s)-1]
	return value
}

func printStack(s *stack) {
	fmt.Print("start$$")
	for _, val := range *s {
		fmt.Print("'")
		fmt.Print(string(val))
		fmt.Print("'")
		fmt.Print(",")
	}
	fmt.Print("end$$\n")
}

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	cost := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	stack := &stack{}
	score := 0
	corrupted := 0

	for s.Scan() {
		skip := false
		for _, val := range s.Text() {
			switch val {
			// opening
			case '(', '[', '{', '<':
				stack.push(val)
			// closing
			case ')', ']', '}', '>':
				if pair[val] != stack.peek() {
					corrupted++
					score += cost[val]
					skip = true
				} else {
					stack.pop()
				}
			}
			if skip {
				break
			}
		}
		if skip {
			continue
		}
	}
	fmt.Println(corrupted)
	fmt.Println(score)
}

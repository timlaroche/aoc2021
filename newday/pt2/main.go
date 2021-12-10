package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type stack []rune

func (s *stack) push(v rune) {
	*s = append(*s, v)
}

func (s *stack) pop() rune {
	if len(*s) == 0 {
		return 'x'
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func (s *stack) peek() rune {
	if len(*s) == 0 {
		return 'x'
	}
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

var cost map[rune]int

var pair map[rune]rune

var closingPair map[rune]rune

var completeScore map[rune]int

func main() {
	f, _ := os.Open("../input copy")
	s := bufio.NewScanner(f)

	// global setup
	cost = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	pair = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	closingPair = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	completeScore = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	// stack := &stack{}
	// score := 0
	corrupted := 0
	allScores := []int{}

	for s.Scan() {
		if !isCorrupted(s.Text()) {
			fmt.Println("got incomplete")
			fmt.Println(isValid(s.Text()))
		}
	}
	fmt.Println(corrupted)
	sort.Ints(allScores)
	fmt.Println(allScores)
}

// skip corrupted
func isCorrupted(s string) bool {
	internalStack := &stack{}
	for _, val := range s {
		switch val {
		// opening
		case '(', '[', '{', '<':
			internalStack.push(val)
		// closing
		case ')', ']', '}', '>':
			if pair[val] != internalStack.peek() {
				return true
			} else {
				internalStack.pop()
			}
		}
	}
	return false
}

func isValid(s string) bool {
	internalStack := &stack{}
	for _, val := range s {
		switch val {
		// opening
		case '(', '[', '{', '<':
			internalStack.push(val)
		// closing
		case ')', ']', '}', '>':
			if pair[val] != internalStack.peek() {
				return false
			} else {
				internalStack.pop()
			}
		}
	}
	return true
}

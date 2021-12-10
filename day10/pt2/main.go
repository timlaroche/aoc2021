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

	closingPair := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	sbScore := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	score := 0
	corrupted := 0
	allScores := []int{}

	for s.Scan() {
		skip := false
		stack := &stack{}
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
		} else {
			sb := ""
			score := 0
			for stack.peek() != 'x' {
				w := stack.pop()
				sb += string(closingPair[w])
			}
			for _, char := range sb {
				score = (score * 5) + sbScore[char]
			}
			fmt.Printf("%s:%d\n", sb, score)
			allScores = append(allScores, score)
		}
	}
	sort.Ints(allScores)
	fmt.Println(len(allScores))
	fmt.Println(allScores)
	fmt.Println(allScores[((len(allScores)+1)/2)-1])
}

// 5487092923 too high

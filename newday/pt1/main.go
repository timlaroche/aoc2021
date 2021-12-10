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
	f, _ := os.Open("../input copy")
	s := bufio.NewScanner(f)

	cost := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	stack := &stack{}
	score := 0
	corrupted := 0

	for s.Scan() {
		early := false
		for _, val := range s.Text() {
			switch val {
			// opening
			case '(', '[', '{', '<':
				stack.push(val)
			// closing
			case ')':
				if stack.peek() == '(' {
					stack.pop()
				} else {
					fmt.Printf("expected %c but found %c\n", stack.peek(), val)
					fmt.Printf("adding %d to score using key\n", cost[')'])
					score += cost[')']
					early = true
					corrupted++
					break
				}
			case ']':
				if stack.peek() == '[' {
					stack.pop()
				} else {
					fmt.Printf("expected %c but found %c\n", stack.peek(), val)
					fmt.Printf("adding %d to score using key\n", cost[']'])
					score += cost[']']
					early = true
					corrupted++
					break
				}
			case '}':
				if stack.peek() == '{' {
					stack.pop()
				} else {
					fmt.Printf("expected %c but found %c\n", stack.peek(), val)
					fmt.Printf("adding %d to score using key \n", cost['}'])

					score += cost['}']
					early = true
					corrupted++

					break
				}
			case '>':
				if stack.peek() == '<' {
					stack.pop()
				} else {
					fmt.Printf("expected %c but found %c\n", stack.peek(), val)
					fmt.Printf("adding %d to score using key \n", cost['>'])
					score += cost['>']
					early = true
					corrupted++

					break
				}
			}
		}
		if early {
		}
	}
	fmt.Println(corrupted)
	fmt.Println(score)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var rules map[string]string

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	//input str
	s.Scan()
	startState := s.Text()
	s.Scan()
	fmt.Println(startState)

	//parse rules
	rules = map[string]string{}
	count := map[rune]int{}
	for s.Scan() {
		rule := regexp.MustCompile("->").Split(s.Text(), -1)
		key := strings.TrimSpace(rule[0])
		val := strings.TrimSpace(rule[1])
		rules[key] = val
	}

	//fmt.Println(rules)
	for i := 0; i < 10; i++ {
		startState = step(startState)
	}
	fmt.Println(startState)
	for _, char := range startState {
		count[char] += 1
	}
	fmt.Println(count)
}

func step(input string) string {
	j := 1
	sb := ""
	for i, _ := range input {
		if i == len(input)-1 {
			sb += fmt.Sprintf("%c", input[i])
			return sb
		}
		fmt.Printf("comparing %c%c\n", input[i], input[j])
		key := fmt.Sprintf("%c%c", input[i], input[j])
		fmt.Println(key)
		if val, ok := rules[key]; ok {
			sb += fmt.Sprintf("%c%s", input[i], val)
		}
		j++
	}
	return sb
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	//input str
	s.Scan()
	startState := s.Text()
	s.Scan()
	fmt.Println(startState)

	//parse rules
	rules := map[string]string{}
	genRules1 := map[string]string{}
	genRules2 := map[string]string{}
	count := map[string]int{}
	for s.Scan() {
		rule := regexp.MustCompile("->").Split(s.Text(), -1)
		key := strings.TrimSpace(rule[0])
		val := strings.TrimSpace(rule[1])
		rules[key] = val
	}

	// generate new rules
	for k, v := range rules {
		rule1 := fmt.Sprintf("%c%s", k[0], v)
		rule2 := fmt.Sprintf("%s%c", v, k[1])

		genRules1[k] = rule1
		genRules2[k] = rule2
	}

	// fmt.Println("====")
	// fmt.Printf("rules: %v\n", rules)
	// fmt.Println("====")
	// fmt.Printf("genRules1: %v\n", genRules1)
	// fmt.Println("====")
	// fmt.Printf("genRules2: %v\n", genRules2)
	// fmt.Println("====")

	//fmt.Println(rules)
	count = initMap(startState)

	fmt.Printf("start count: %v\n", count)

	for i := 0; i < 40; i++ {
		count = step(count, genRules1, genRules2)
	}

	finalCount := map[string]int{}
	for k, v := range count {
		finalCount[string(k[0])] += v
		finalCount[string(k[1])] += v
	}
	//fmt.Println(count)
	fmt.Println(finalCount)

	counts := []int{}
	for _, v := range finalCount {
		counts = append(counts, int(math.Ceil(float64(v/2))))
	}
	sort.Ints(counts) // off by one error.... manually accounted for
	fmt.Println(counts)
}

func initMap(input string) map[string]int {
	j := 1
	m := map[string]int{}
	for i := range input {
		if i == len(input)-1 {
			return m
		}
		key := fmt.Sprintf("%c%c", input[i], input[j])
		m[key] += 1
		j++
	}
	return m
}

func step(input map[string]int, genRules1 map[string]string, genRules2 map[string]string) map[string]int {
	m := map[string]int{}
	for k, v := range input {
		r1 := genRules1[k]
		r2 := genRules2[k]
		//fmt.Printf("k: %s, r1: %s, r2:%s\n", k, r1, r2)
		m[r1] += v
		m[r2] += v
	}

	fmt.Println("==after step==")
	fmt.Println(m)
	return m
}

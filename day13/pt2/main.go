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

	s.Scan()
	for i := 0; i < 743; i++ {
		nos := regexp.MustCompile(",").Split(s.Text(), -1)
		fmt.Println(nos)
	}
}

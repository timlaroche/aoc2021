package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var flashes int // lol global

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	state := [10][10]int{}
	i := 0
	for s.Scan() {
		for j, n := range s.Text() {
			intVal, _ := strconv.Atoi(string(n))
			fmt.Println(intVal)
			state[i][j] = intVal
		}
		i++
	}

	state = step(state)
	fmt.Println(state)
}

func step(board [10][10]int) [10][10]int {
	//step 1
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board[i][j] += 1
		}
	}
	//step 2

	return board
}

func checkFlash(board [10][10]int) [10][10]int {
	shouldLoop := false

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j] >= 9 {
				u := board
				d :=
				l :=
				r :=
				ul :=
				ur :=
				dl := 
				dr :=
			}
		}
	}
}

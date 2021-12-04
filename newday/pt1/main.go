package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var s *bufio.Scanner

func main() {
	f, _ := os.Open("../input")
	s = bufio.NewScanner(f)
	s.Scan() // first line
	reg := regexp.MustCompile(",")
	winning := s.Text()
	winningNumbers := reg.Split(winning, -1)
	winningIntArray := []int{}
	winningMap := map[int]bool{}

	for _, val := range winningNumbers {
		intVal, _ := strconv.Atoi(val)
		winningIntArray = append(winningIntArray, intVal)
	}

	s.Scan() // whiteline

	found := false
	for _, val := range winningIntArray {
		winningMap[val] = true
		// start from the top
		f, _ := os.Open("../input")
		s = bufio.NewScanner(f)
		s.Scan()
		s.Scan()
		// naughty
		for b := 0; b < 80; b++ {
			board := genBoard(s)
			fmt.Println("=====")
			fmt.Println(board)
			fmt.Println("=====")
			if checkWin(board, winningMap) {
				fmt.Println("winner")
				found = true
				justCalled := val
				sum := sumBoardAgainstWinning(board, winningMap)
				fmt.Println(justCalled)
				fmt.Println(sum)
				fmt.Println(sum * justCalled)
				break
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
}

func sumBoardAgainstWinning(board [5][5]int, winningMap map[int]bool) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := board[i][j]
			if !winningMap[val] {
				sum += val
			} else {
				fmt.Println(val)
			}
		}
	}
	return sum
}

func genBoard(s *bufio.Scanner) [5][5]int {
	grid := [5][5]int{}
	//reg := regexp.MustCompile(` \s?\d+`)
	for i := 0; i < 5; i++ {
		s.Scan()
		//nums := reg.Split(s.Text(), -1)
		nums := strings.Split(s.Text(), " ")
		//fmt.Println(nums)
		workingNums := removeEmptyStrings(nums)
		for j := 0; j < 5; j++ {
			val, _ := strconv.Atoi(strings.TrimSpace(workingNums[j]))
			grid[i][j] = val
		}
	}
	s.Scan()
	return grid
}

func checkWin(grid [5][5]int, winning map[int]bool) bool {
	// check horizontal
	for i := 0; i < 5; i++ {
		horCount := 0
		for j := 0; j < 5; j++ {
			if winning[grid[i][j]] {
				horCount++
			} else {
				horCount = 0
			}
		}
		if horCount == 5 {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		vertCount := 0
		for j := 0; j < 5; j++ {
			if winning[grid[j][i]] {
				vertCount++
			} else {
				vertCount = 0
			}
		}
		if vertCount == 5 {
			return true
		}
	}

	return false
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

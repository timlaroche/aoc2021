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

type holding struct {
	value int
	mark  bool
}

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

	boards := make([][5][5]*holding, 100)
	winningBoards := map[int]bool{}
	//winCount := 0
	for b := 0; b < 100; b++ {
		boards[b] = genBoard(s)
	}

	for _, val := range winningIntArray {
		winningMap[val] = true
		for i := 0; i < 100; i++ {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if boards[i][j][k].value == val {
						boards[i][j][k].mark = true
					}
				}
			}
			shouldRemove := shouldRemove(boards[i])

			if shouldRemove {
				winningBoards[i] = true
			}
			if len(winningBoards) == 99 {
				fmt.Println(i)
				fmt.Println(sumBoardAgainstWinning(boards[66], winningMap))
				fmt.Println(val)
				for i, val := range winningBoards {
					if winningBoards[i] != true {
						fmt.Println(i)
						fmt.Println(val)
					}
				}
				return
			}
			// 	i
			// 	if winningBoards[i] == true {
			// 		winCount--
			// 	}
			// 	winningBoards[i] = true
			// }
			// if winCount == 99 {
			// 	fmt.Println(i)
			// 	fmt.Println(val)
			// 	fmt.Println(sumBoardAgainstWinning(boards[i], winningMap))
			// 	printBoard(boards[i])
			// 	//fmt.Println(val * sumBoardAgainstWinning(boards[i], winningMap))
			// 	return
			// }
		}
	}
}

func printBoard(board [5][5]*holding) {
	for i := 0; i < 5; i++ {
		fmt.Println()
		for j := 0; j < 5; j++ {
			fmt.Printf(" %d", board[i][j].value)
		}
	}
}

func shouldRemove(board [5][5]*holding) bool {
	// check horizontal
	for i := 0; i < 5; i++ {
		horCount := 0
		for j := 0; j < 5; j++ {
			if board[i][j].mark == true {
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
			if board[j][i].mark == true {
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

func sumBoardAgainstWinning(board [5][5]*holding, winningMap map[int]bool) int {
	sum := 0
	//fmt.Print(board)
	for i := 0; i < 5; i++ {
		horCount := 0
		for j := 0; j < 5; j++ {
			if winningMap[board[i][j].value] {
				horCount++
			} else {
				horCount = 0
			}
			if horCount == 5 {
				board[i][0].value = 0
				board[i][1].value = 0
				board[i][2].value = 0
				board[i][3].value = 0
				board[i][4].value = 0
				break
			}
		}
		if horCount == 5 {
			break
		}
	}
	for i := 0; i < 5; i++ {
		vertCount := 0
		for j := 0; j < 5; j++ {
			if winningMap[board[j][i].value] {
				vertCount++
			} else {
				vertCount = 0
			}
			if vertCount == 5 {
				board[0][j].value = 0
				board[1][j].value = 0
				board[2][j].value = 0
				board[3][j].value = 0
				board[4][j].value = 0
			}
		}
		if vertCount == 5 {
			break
		}
	}
	// fmt.Println("updated board")
	// fmt.Println(board)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			val := board[i][j].value
			sum += val
		}
	}
	return sum
}

func genBoard(s *bufio.Scanner) [5][5]*holding {
	grid := [5][5]*holding{}
	//reg := regexp.MustCompile(` \s?\d+`)
	for i := 0; i < 5; i++ {
		s.Scan()
		//nums := reg.Split(s.Text(), -1)
		nums := strings.Split(s.Text(), " ")
		//fmt.Println(nums)
		workingNums := removeEmptyStrings(nums)
		for j := 0; j < 5; j++ {
			val, _ := strconv.Atoi(strings.TrimSpace(workingNums[j]))
			grid[i][j] = &holding{
				value: val,
				mark:  false,
			}
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
			if grid[i][j] == 0 {
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

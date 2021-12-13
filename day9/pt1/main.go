package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	grid := [102][102]int{}

	// Pad grid
	for j := 0; j < 102; j++ {
		grid[0][j] = 9999
	}
	for j := 0; j < 102; j++ {
		grid[101][j] = 9999
	}
	for i := 0; i < 102; i++ {
		grid[i][0] = 9999
	}
	for i := 0; i < 102; i++ {
		grid[i][101] = 9999
	}

	// Fill in inside the padding
	i := 1
	for s.Scan() {
		j := 1
		for _, x := range s.Text() {
			intVal, _ := strconv.Atoi(string(x))
			grid[i][j] = intVal
			j++
		}
		i++
	}

	// check all around
	sum := 0
	for i := 1; i < 101; i++ {
		for j := 1; j < 101; j++ {
			// up_i, up_j := i-1, j
			// down_i, down_j := i+1, j
			// left_i, left_j := i, j-1
			// right_i, right_j := i, j+1

			up := grid[i-1][j]
			down := grid[i+1][j]
			left := grid[i][j-1]
			right := grid[i][j+1]
			current := grid[i][j]

			if current < up && current < down && current < left && current < right {
				sum += current
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

// 11828 too high

// func main() {
// 	f, _ := os.Open("../input copy")
// 	s := bufio.NewScanner(f)

// 	grid := [7][12]int{}

// 	// Pad grid
// 	for j := 0; j < 12; j++ {
// 		grid[0][j] = 9999
// 	}
// 	for j := 0; j < 12; j++ {
// 		grid[6][j] = 9999
// 	}
// 	for i := 0; i < 7; i++ {
// 		grid[i][0] = 9999
// 	}
// 	for i := 0; i < 7; i++ {
// 		grid[i][11] = 9999
// 	}

// 	//fmt.Println(grid)

// 	// Fill in inside the padding
// 	i := 1
// 	for s.Scan() {
// 		j := 1
// 		for _, x := range s.Text() {
// 			intVal, _ := strconv.Atoi(string(x))
// 			grid[i][j] = intVal
// 			//fmt.Printf("i: %d, j: %d, val:%d \n", i, j, grid[i][j])
// 			j++
// 		}
// 		i++
// 	}
// 	fmt.Println(grid)
// 	// check all around
// 	sum := 0
// 	for i := 1; i < 6; i++ {
// 		for j := 1; j < 11; j++ {
// 			// up_i, up_j := i-1, j
// 			// down_i, down_j := i+1, j
// 			// left_i, left_j := i, j-1
// 			// right_i, right_j := i, j+1

// 			up := grid[i-1][j]
// 			down := grid[i+1][j]
// 			left := grid[i][j-1]
// 			right := grid[i][j+1]
// 			current := grid[i][j]

// 			fmt.Printf("comparing: %d \n", current)
// 			if current < up && current < down && current < left && current < right {
// 				fmt.Println(current)
// 				sum += (current + 1)
// 			}
// 		}
// 		fmt.Println("===")
// 	}

// 	fmt.Println(sum)
// }

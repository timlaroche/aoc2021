package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x   int
	y   int
	val int
}

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

	// get basin points
	basins := []Point{}
	sum := 0
	for i := 1; i < 101; i++ {
		for j := 1; j < 101; j++ {
			if checkAllAdjacent(i, j, grid, true) {
				basins = append(basins, *&Point{i, j, grid[i][j]})
				sum += grid[i][j] + 1
			}
		}
	}
	basinScore := make([]int, len(basins))
	fmt.Println(sum)

	// first loop
	for i, point := range basins {
		queue := []Point{}
		basinScore[i] += 1
		if checkAllAdjacent(point.x, point.y, grid, true) {
			lp, _ := checkLeft(point.x, point.y, grid)
			rp, _ := checkRight(point.x, point.y, grid)
			up, _ := checkUp(point.x, point.y, grid)
			dp, _ := checkDown(point.x, point.y, grid)
			queue = append(queue, lp, rp, up, dp)
		}
		fmt.Println(queue)
		queue = removeNines(queue)
		fmt.Println(queue)
		//now loop onwards until empty
		for len(queue) != 0 {
			fmt.Println(queue)
			for _, point := range queue {
				queue = queue[1:]
				basinScore[i] += 1
				if checkAllAdjacent(point.x, point.y, grid, false) {
					lp, _ := checkLeft(point.x, point.y, grid)
					rp, _ := checkRight(point.x, point.y, grid)
					up, _ := checkUp(point.x, point.y, grid)
					dp, _ := checkDown(point.x, point.y, grid)
					queue = append(queue, lp, rp, up, dp)
				}
				queue = removeNines(queue)
			}
		}
	}

	fmt.Println(basinScore)
}

func removeNines(b []Point) []Point {
	var newB []Point
	for _, point := range b {
		if point.val < 9 {
			newB = append(newB, point)
		}
	}
	return newB
}

func outOfBounds(i int, j int, grid [102][102]int) bool {
	defer func() bool {
		if recover() != nil {
			return false
		}
		return true
	}()
	_ = grid[i][j]
	return true
}

func checkAllAdjacent(i int, j int, grid [102][102]int, desc bool) bool {
	u := grid[i-1][j]
	d := grid[i+1][j]
	l := grid[i][j-1]
	r := grid[i][j+1]
	current := grid[i][j]

	if desc {
		if current < u && current < d && current < l && current < r {
			return true
		}
		return false
	} else {
		if current > u && current > d && current > l && current > r {
			return true
		}
		return false
	}
}

func checkLeft(i int, j int, grid [102][102]int) (Point, bool) {
	left := grid[i][j-1]
	//fmt.Printf("left: %d vs  current: %d\n", left, grid[i][j])
	if left > grid[i][j] {
		return *&Point{i, j - 1, left}, true
	}
	return *&Point{9999, 9999, 9999}, false
}
func checkRight(i int, j int, grid [102][102]int) (Point, bool) {
	right := grid[i][j+1]
	if right > grid[i][j] {
		return *&Point{i, j + 1, right}, true
	}
	return *&Point{9999, 9999, 9999}, false
}
func checkUp(i int, j int, grid [102][102]int) (Point, bool) {
	up := grid[i-1][j]
	if up > grid[i][j] {
		return *&Point{i - 1, j, up}, true
	}
	return *&Point{9999, 9999, 9999}, false
}
func checkDown(i int, j int, grid [102][102]int) (Point, bool) {
	down := grid[i+1][j]
	if down > grid[i][j] {
		return *&Point{i + 1, j, down}, true
	}
	return *&Point{9999, 9999, 9999}, false
}

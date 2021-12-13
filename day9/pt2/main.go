package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for j := 0; j < 12; j++ {
		grid[0][j] = 9999
	}
	for j := 0; j < 12; j++ {
		grid[101][j] = 9999
	}
	for i := 0; i < 7; i++ {
		grid[i][0] = 9999
	}
	for i := 0; i < 7; i++ {
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

	fmt.Println(grid)

	// TEST
	somePoints := make([]Point, 2)
	somePoints[0] = *&Point{1, 1, 500}
	somePoints[1] = *&Point{1, 1, 500}
	fmt.Println(somePoints)
	somePoints = removeDuplicates(somePoints)
	fmt.Println(somePoints)
	//

	// get basin points
	basins := []Point{}
	sum := 0
	for i := 1; i < 101; i++ {
		for j := 1; j < 101; j++ {
			if checkAllAdjacent(i, j, grid, true) {
				basins = append(basins, *&Point{i, j, grid[i][j]})
				sum += (grid[i][j] + 1)
			}
		}
	}
	basinScore := make([]int, len(basins))
	fmt.Println(sum)
	fmt.Println(basins)
	fmt.Println(grid[1][1])
	fmt.Println("====")
	// first loop
	for i, point := range basins {
		visited := map[string]bool{}
		queue := []Point{}
		basinScore[i] += 1
		lp, _ := checkLeft(point.x, point.y, grid)
		rp, _ := checkRight(point.x, point.y, grid)
		up, _ := checkUp(point.x, point.y, grid)
		dp, _ := checkDown(point.x, point.y, grid)
		fmt.Printf("Checking point: x:%d, y:%d, val:%d\n", point.x, point.y, point.val)
		fmt.Printf("u: %d,d: %d,l: %d,r: %d\n", up.val, dp.val, lp.val, rp.val)
		queue = append(queue, lp, rp, up, dp)
		queue = removeNines(queue)
		queue = removeDuplicates(queue)
		//now loop onwards until empty
		for len(queue) != 0 {
			basinScore[i] += 1
			fmt.Println(queue)
			queuePoint := queue[0]
			visited[fmt.Sprintf("%d%d", queuePoint.x, queuePoint.y)] = true
			queue = queue[1:]
			fmt.Printf("Checking point: x:%d, y:%d, val:%d\n", queuePoint.x, queuePoint.y, queuePoint.val)
			lp, _ := checkLeft(queuePoint.x, queuePoint.y, grid)
			rp, _ := checkRight(queuePoint.x, queuePoint.y, grid)
			up, _ := checkUp(queuePoint.x, queuePoint.y, grid)
			dp, _ := checkDown(queuePoint.x, queuePoint.y, grid)
			fmt.Printf("u: %d,d: %d,l: %d,r: %d\n", up.val, dp.val, lp.val, rp.val)
			// dont go backwards
			switch queuePoint.val {
			case lp.val:
				fmt.Println("dont go back left")
				lp.val = 99998
			case rp.val:
				fmt.Println("dont go back right")
				rp.val = 99998
			case up.val:
				fmt.Println("dont go back up")
				up.val = 99998
			case dp.val:
				fmt.Println("dont go back down")
				dp.val = 99998
			}
			if !visited[fmt.Sprintf("%d%d", lp.x, lp.y)] {
				queue = append(queue, lp)

			}
			if !visited[fmt.Sprintf("%d%d", rp.x, rp.y)] {
				queue = append(queue, rp)

			}
			if !visited[fmt.Sprintf("%d%d", dp.x, dp.y)] {
				queue = append(queue, dp)

			}
			if !visited[fmt.Sprintf("%d%d", up.x, up.y)] {
				queue = append(queue, up)
			}
			queue = removeNines(queue)
			queue = removeDuplicates(queue)
			fmt.Printf("Updated queue: %v\n", queue)
		}
		fmt.Println(basinScore[i])
		fmt.Println("===newbasin===")
	}

	sort.Ints(basinScore)
	fmt.Println(basinScore)
	size := len(basinScore)
	finalScore := basinScore[size-1] * basinScore[size-2] * basinScore[size-3]
	fmt.Println(finalScore)
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

func removeDuplicates(b []Point) []Point {
	newB := []Point{}
	exists := make(map[string]bool)
	for _, point := range b {
		key := fmt.Sprintf("%d%d", point.x, point.y)
		if _, exist := exists[key]; !exist {
			exists[key] = true
			newB = append(newB, point)
		}
	}
	//fmt.Println(exists)
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
	fmt.Printf("u:%d d:%d l:%d r:%d current:%d \n", u, d, l, r, current)
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

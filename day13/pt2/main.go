package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	points := map[Point]bool{}

	s.Scan()
	for i := 0; i < 743; i++ {
		nos := regexp.MustCompile(",").Split(s.Text(), -1)
		//fmt.Println(nos)
		x, _ := strconv.Atoi(nos[0])
		y, _ := strconv.Atoi(nos[1])
		//fmt.Printf("%d, %d\n", x, y)
		point := *&Point{x, y}
		points[point] = true
		s.Scan()
	}

	// fmt.Println(points)

	for i := 0; i < 12; i++ {
		fmt.Println("========")
		s.Scan()
		fmt.Println(s.Text())
		losString := regexp.MustCompile("=").Split(s.Text(), -1)
		endex := len(losString) - 1 // end index
		axisVal, _ := strconv.Atoi(losString[endex])
		axis := losString[endex-1]

		fmt.Println(axis)
		if axis == "fold along y" {
			points = transformGivenLine('y', axisVal, points)
			fmt.Println(len(points))
		} else if axis == "fold along x" {
			points = transformGivenLine('x', axisVal, points)
			fmt.Println(len(points))
		}
		// fmt.Print("x[")
		// for k := range points {
		// 	fmt.Printf("%d,", k.x)
		// }
		// fmt.Print("\n")
		// fmt.Print("y[")
		// for k := range points {
		// 	fmt.Printf("%d,", k.y)
		// }
		// // fmt.Print("]")
		// // fmt.Print("\n")
		// // fmt.Println("=======")
	}

	prettyPrint(points)
}

func prettyPrint(t map[Point]bool) {
	size := 50
	internal := [50][50]string{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if t[*&Point{j, i}] {
				internal[j][i] = " #"
			} else {
				internal[j][i] = "  "
			}
		}
	}

	for i := 0; i < size; i++ {
		sb := ""
		for j := 0; j < size+1; j++ {
			if j == size {
				sb += "\n"
			} else {
				sb += internal[j][i]
			}
		}
		fmt.Println(sb)
	}
}

func transformGivenLine(axis rune, axisVal int, points map[Point]bool) map[Point]bool {
	// fold up
	if axis == 'y' {
		fmt.Printf("working on y=%d\n", axisVal)
		// separate points
		set1 := map[Point]bool{} // all points above the line
		set2 := map[Point]bool{} // all points below the line

		for i := 0; i < 2000; i++ {
			for j := 0; j < axisVal; j++ {
				key := *&Point{i, j}
				if points[key] {
					set1[key] = true
				}
			}
		}

		for i := 0; i < 2000; i++ {
			for j := axisVal + 1; j < 2000; j++ {
				key := *&Point{i, j}
				if points[key] {
					set2[key] = true
				}
			}
		}
		// fmt.Println(set1)
		// fmt.Println(len(set1))
		// fmt.Println(set2)
		// fmt.Println(len(set2))

		// transform set 2
		transformedSet2 := map[Point]bool{}
		for k := range set2 {
			difference := k.y - axisVal
			newY := k.y - difference - difference
			//fmt.Printf("key: %v, difference: %d, newY: %d\n", k, difference, newY)
			newPoint := *&Point{k.x, newY}
			transformedSet2[newPoint] = true
		}

		// union transformed set 2 + set 1
		union := map[Point]bool{}
		for k := range set1 {
			union[k] = true
		}
		for k := range transformedSet2 {
			union[k] = true
		}
		return union
	} else { // fold left
		fmt.Printf("working on x=%d\n", axisVal)
		// separate points
		set1 := map[Point]bool{} // all points above the line
		set2 := map[Point]bool{} // all points below the line

		for i := 0; i < axisVal; i++ {
			for j := 0; j < 2000; j++ {
				key := *&Point{i, j}
				if points[key] {
					set1[key] = true
				}
			}
		}

		for i := axisVal + 1; i < 2000; i++ {
			for j := 0; j < 2000; j++ {
				key := *&Point{i, j}
				if points[key] {
					set2[key] = true
				}
			}
		}
		// fmt.Println(set1)
		// fmt.Println(len(set1))
		// fmt.Println(set2)
		// fmt.Println(len(set2))

		// transform set 2
		transformedSet2 := map[Point]bool{}
		for k := range set2 {
			difference := k.x - axisVal
			newX := k.x - difference - difference
			//fmt.Printf("key: %v, difference: %d, newY: %d\n", k, difference, newX)
			newPoint := *&Point{newX, k.y}
			transformedSet2[newPoint] = true
		}

		// union transformed set 2 + set 1
		union := map[Point]bool{}
		for k := range set1 {
			union[k] = true
		}
		for k := range transformedSet2 {
			union[k] = true
		}
		return union
	}
}

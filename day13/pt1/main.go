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

	s.Scan()
	// fmt.Println(s.Text())
	losString := regexp.MustCompile("=").Split(s.Text(), -1)
	endex := len(losString) - 1 // end index
	axisVal, _ := strconv.Atoi(losString[endex])
	axis := losString[endex-1]

	fmt.Println(axis)
	if axis == "fold along y" {
		new := transformGivenLine('y', axisVal, points)
		fmt.Println("-===-")
		//fmt.Println(new)
		fmt.Println(len(new))
	} else if axis == "fold along x" {
		new := transformGivenLine('x', axisVal, points)
		fmt.Println("-===-")
		//fmt.Println(new)
		fmt.Println(len(new))
	}
}

func transformGivenLine(axis rune, axisVal int, points map[Point]bool) map[Point]bool {
	// fold up
	if axis == 'y' {
		fmt.Println("working on y reflection")
		fmt.Println(axisVal)
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
		fmt.Println("working on x reflection")
		fmt.Println(axisVal)
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
	// 743 too high
}

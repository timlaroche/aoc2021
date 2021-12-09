package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("../input")
	s := bufio.NewScanner(f)

	grid := [1000][1000]int{}

	for s.Scan() {
		arrowRegex := regexp.MustCompile("->")
		commaRegex := regexp.MustCompile(",")
		coordPair := arrowRegex.Split(s.Text(), -1)

		pair1 := commaRegex.Split(coordPair[0], -1)
		p1x, _ := strconv.Atoi(strings.TrimSpace(pair1[0]))
		p1y, _ := strconv.Atoi(strings.TrimSpace(pair1[1]))
		pair2 := commaRegex.Split(coordPair[1], -1)
		p2x, _ := strconv.Atoi(strings.TrimSpace(pair2[0]))
		p2y, _ := strconv.Atoi(strings.TrimSpace(pair2[1]))

		dy := p2y - p1y
		dx := p2x - p1x
		beZero := false
		if dy == 0 || dx == 0 {
			beZero = true
		}
		var grad int
		if beZero {
			grad = 0
		} else {
			grad = dy / dx
		}

		if grad == 1 || grad == -1 {
			counterx := p1x
			countery := p1y
			if dy > 0 && dx > 0 {
				for i := 0; i < dy+1; i++ {
					grid[counterx+i][countery+i] += 1
				}
				continue
			} else if dy > 0 && dx < 0 {
				for i := 0; i < dy+1; i++ {
					grid[counterx-i][countery+i] += 1
				}
				continue
			} else if dy < 0 && dx > 0 {
				for i := 0; i < dx+1; i++ {
					grid[counterx+i][countery-i] += 1
				}
				continue
			} else if dy < 0 && dx < 0 {
				for i := 0; i < (dy*-1)+1; i++ {
					grid[counterx-i][countery-i] += 1
				}
				continue
			}
		}

		if p1x == p2x {
			bigger := 0
			smaller := 0
			if p1y < p2y {
				bigger = p2y
				smaller = p1y
			} else if p1y > p2y {
				bigger = p1y
				smaller = p2y
			} else if p1y == p2y {
				grid[p1y][p1x] += 1
				continue
			}
			for i := smaller; i < bigger+1; i++ {
				grid[p1x][i] += 1
			}
		}

		if p1y == p2y {
			bigger := 0
			smaller := 0
			if p1x < p2x {
				bigger = p2x
				smaller = p1x
			} else if p1x > p2x {
				bigger = p1x
				smaller = p2x
			} else if p1x == p2x {
				grid[p1y][p1x] += 1
				continue
			}
			for i := smaller; i < bigger+1; i++ {
				grid[i][p1y] += 1
			}
		}

	}

	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] >= 2 {
				count += 1
			}
		}
	}
	fmt.Println("--- count ---")
	fmt.Println(count)
}

// 8239 too low
// 8240 too low
// 16436 too low
// 16469 wrong
// 21589 wrong
// 21113 wrong

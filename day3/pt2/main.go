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

	oxyList := make([]string, 1000)
	co2List := make([]string, 1000)

	i := 0

	// Load values
	for s.Scan() {
		oxyList[i] = s.Text()
		co2List[i] = s.Text()
		i++
	}

	for i := 0; i < 12; i++ {
		one := oneMostCommon(i, oxyList, false)
		if one {
			for j, value := range oxyList {
				if value[i] != '1' {
					oxyList[j] = ""
				}
			}
		} else {
			for j, value := range oxyList {
				if value[i] != '0' {
					oxyList[j] = ""

				}
			}
		}
		removeEmptyStrings(&oxyList)
		if len(oxyList) == 1 {
			break
		}
	}

	for i := 0; i < 12; i++ {
		one := oneLeastCommon(i, co2List, true)
		fmt.Println(one)
		if one {
			for j, value := range co2List {
				if value[i] != '1' {
					co2List[j] = ""
				}
			}
		} else {
			for j, value := range co2List {
				if value[i] != '0' {
					co2List[j] = ""

				}
			}
		}
		removeEmptyStrings(&co2List)
		if len(co2List) == 1 {
			break
		}
	}

	fmt.Println(oxyList)
	fmt.Println(co2List)

	oxyDecimal, _ := strconv.ParseInt(oxyList[0], 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2List[0], 2, 64)

	fmt.Println(oxyDecimal * co2Decimal)
}

func oneMostCommon(index int, numbers []string, takeZero bool) bool {
	oneCount := 0
	zeroCount := 0
	for _, number := range numbers {
		//fmt.Println(number)
		if number[index] == '1' {
			oneCount += 1
		} else {
			zeroCount += 1
		}
	}
	if oneCount == zeroCount {
		if takeZero {
			return false
		} else {
			return true
		}
	}
	if oneCount > zeroCount {
		return true
	}
	return false
}

func oneLeastCommon(index int, numbers []string, takeZero bool) bool {
	oneCount := 0
	zeroCount := 0
	for _, number := range numbers {
		//fmt.Println(number)
		if number[index] == '1' {
			oneCount += 1
		} else {
			zeroCount += 1
		}
	}
	if oneCount == zeroCount {
		if takeZero {
			return false
		} else {
			return true
		}
	}
	if oneCount < zeroCount {
		return true
	}
	return false
}

func removeEmptyStrings(s *[]string) {
	var r []string
	for _, str := range *s {
		if str != "" {
			r = append(r, str)
		}
	}
	*s = r
}

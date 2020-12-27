package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("thirteen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := -1
	bestDepartureTime := 1000000000
	ans2 := 0

	line, _, err := reader.ReadLine()
	startTime, err := strconv.Atoi(string(line))
	if err != nil {
		panic("unable to parse start time")
	}
	line, _, err = reader.ReadLine()

	buses := strings.Split(string(line), ",")

	mul := 1
	for i, b := range buses {
		if b == "x" {
			continue
		}
		num, err := strconv.Atoi(b)
		if err != nil {
			panic("unable to parse bus number")
		}
		count := startTime / num
		departureTime := count * num
		if departureTime < startTime {
			departureTime += num
		}
		if departureTime < bestDepartureTime {
			bestDepartureTime = departureTime
			ans1 = (bestDepartureTime - startTime) * num
		}
		target := num - i
		for target >= num {
			target -= num
		}
		for target < 0 {
			target += num
		}

		prev := ans2
		for ans2 % num != target {
			ans2 += mul
			if ans2 < prev {
				println("err")
				break
			}
		}
		mul *= num
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	input, err := os.Open("five/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 0
	allSeats := make([]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)
		seatId := getSeatId(sLine)
		allSeats = append(allSeats, seatId)
		if seatId > ans1 {
			ans1 = seatId
		}
	}
	sort.Ints(allSeats)
	prev := allSeats[0] - 1
	for _,x := range allSeats{
		if x - prev != 1 {
			ans2 = x - 1
		}
		prev = x
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func getSeatId(s string)int{
	row := 0
	col := 0
	dRow := 64
	dCol := 4
	for _,x := range s {
		switch x {
		case'L':
			dCol /= 2
		case'R':
			col += dCol
			dCol /= 2
		case 'F':
			dRow /=2
		case 'B':
			row += dRow
			dRow /=2

		}
	}
	return row * 8 + col
}


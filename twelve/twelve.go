package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	dx := []int {0, 1, 0, -1}
	dy := []int {1, 0, -1, 0}
	dd := []rune {'N', 'E', 'S', 'W'}
	dir := 1
	x, y := 0, 0
	wx, wy := 10, 1
	x2, y2 := 0, 0

	input, err := os.Open("twelve/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	ans1 := 0
	ans2 := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		if len(line)==0{
			break
		}
		sLine := string(line)
		move := rune(sLine[0])
		num, err := strconv.Atoi(sLine[1:])
		if err != nil {
			panic("Unable to parse number")
		}
		switch move {
		case 'L': {
			dir -= num / 90
			for dir < 0 {
				dir += 4
			}
			if num == 180 {
				wx, wy = -wx, -wy
			} else if num == 90 {
				wx, wy = -wy, wx
			} else {
				wx, wy = wy, -wx
			}

		}
		case 'R': {
			dir += num / 90
			for dir > 3 {
				dir -= 4
			}
			if num == 180 {
				wx, wy = -wx, -wy
			} else if num == 90 {
				wx, wy = wy, -wx
			} else {
				wx, wy = -wy, wx
			}
		}

		case 'F': {
			x += dx[dir] * num
			y += dy[dir] * num

			x2 += wx * num
			y2 += wy * num
		}

		default: {
			dirNow := -1
			for d := range dd {
				if dd[d] == move {
					dirNow = d
				}
			}
			x += dx[dirNow] * num
			y += dy[dirNow] * num

			wx += dx[dirNow] * num
			wy += dy[dirNow] * num
		}

		}
	}
	ans1 = int(math.Abs(float64(x)) + math.Abs(float64(y)))
	ans2 = int(math.Abs(float64(x2)) + math.Abs(float64(y2)))

	fmt.Println(ans1)
	fmt.Println(ans2)
}

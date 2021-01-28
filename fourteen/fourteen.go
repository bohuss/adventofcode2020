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

	input, err := os.Open("fourteen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	ans1 := 0
	ans2 := 0

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	mem := make(map[int] int)
	mem2 := make(map[int] int)

	re := regexp.MustCompile("[ =\\[\\]]+")

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)

		if strings.HasPrefix(sLine, "mask"){
			mask = strings.Split(sLine, " ")[2]
		} else {
			a := re.Split(sLine, -1)
			pos, err := strconv.Atoi(a[1])
			if err != nil {
				panic("unable to parse mem position")
			}
			val, err := strconv.Atoi(a[2])
			if err != nil {
				panic("unable to parse value")
			}
			val1 := maskedValue(val, mask)
			mem[pos] = val1

			for _, pos2 := range maskedPositions(pos, mask) {
				mem2[pos2] = val
			}
		}
	}

	for _, val := range mem {
		ans1 += val
	}

	for _, val2 := range mem2 {
		ans2 += val2
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func maskedValue(x int, mask string) int {
	ans := x
	for i := range mask {
		switch mask[i] {
		case '1': {
			ans |= 1 << (len(mask) - i - 1)
		}
		case '0': {
			ans &^= 1 << (len(mask) - i - 1)
		}
		}
	}
	return ans
}

func maskedPositions(pos int, mask string) []int {
	ans := make([]int, 0)
	ans = append(ans, 0)

	for i, v := range mask {
		now := make([]int, 0)
		for _, a := range ans {
			switch v {
			case '0':
				{
					add := 0
					if pos & (1 << (len(mask) - i - 1)) > 0 {
						add = 1
					}
					now = append(now, a * 2 | add)
				}
			case '1':
				{
					now = append(now, a * 2 | 1)
				}
			case 'X':
				{
					now = append(now, a * 2)
					now = append(now, a * 2 | 1)
				}
			}
		}
		ans = now
	}
	return ans
}
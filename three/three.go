package main

import (
	"fmt"
	"os"
)

func main() {

	input, err := os.Open("three/input.txt")

	if err != nil {
		panic(err)
	}

	var line string
	lines := make([]string, 0)

	for {
		_, err := fmt.Fscanf(input, "%s", &line)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		lines = append(lines, line)
	}

	ans1 := count(lines, 3, 1)

	rs := []int{1,1,1,1,2}
	cs := []int{1,3,5,7,1}

	ans2 := 1

	for id := range rs {
		ans2 *= count(lines, cs[id], rs[id])
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func count(lines [] string, dc, dr int) int {
	R := len(lines)
	C := len(lines[0])

	ans := 0
	c := 0
	for i:=0; i<R; i+=dr {
		if lines[i][c] == '#' {
			ans++
		}
		c += dc
		c %= C
	}
	return ans
}


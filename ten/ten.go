package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	input, err := os.Open("ten/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 0
	a := make([]int, 0)
	a = append(a, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		now, err := strconv.Atoi(string(line))
		if err != nil {
			panic("unable to parse number")
		}
		a = append(a, now)

	}
	sort.Ints(a)
	a = append(a, a[len(a) - 1] + 3)
	d := make([]int, 4)
	for i:=1; i<len(a); i++ {
		d[a[i] - a[i - 1]]++
	}

	ans1 = d[1] * d[3]

	n := len(a)
	p := make([]int, n)
	p[0] = 1

	for i:=1; i<len(a); i++ {
		p[i] = p[i - 1]
		for j := 2; i - j >= 0 && a[i - j] >= a[i] - 3; j++ {
			p[i] += p[i - j]
		}
	}
	ans2 = p[len(p) - 1]

	fmt.Println(ans1)
	fmt.Println(ans2)
}

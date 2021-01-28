package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input, err := os.Open("nine/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 0

	a := make([]int, 0)
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

	for id := 25; id < len(a); id++ {
		ok := false
		for i := id - 25; i < id && !ok; i ++ {
			for j := i + 1; j < id; j++ {
				if a[i] + a[j] == a[id] {
					ok = true
					break
				}
			}
		}
		if !ok {
			ans1 = a[id]
			break
		}
	}

	from := 0
	to := -1
	sum := 0

	for sum != ans1 {
		if sum < ans1 {
			to++
			sum += a[to]
		} else {
			sum -= a[from]
			from++
		}
	}
	min := 1<<30
	max := 0
	for id := from; id <= to; id++ {
		if a[id] < min {
			min = a[id]
		}
		if a[id] > max {
			max = a[id]
		}
	}
	ans2 = min + max

	fmt.Println(ans1)
	fmt.Println(ans2)
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("fifteen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	ans1 := 0
	ans2 := 0

	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	a := make ([]int, 0)
	for _, v := range strings.Split(string(line), ",") {
		vv, err := strconv.Atoi(v)
		if err != nil {
			panic("unable to parse number")
		}
		a = append(a, vv)
	}

	lastSeen := make(map[int] int)
	for i, x := range a {
		if i == len(a) - 1 {
			break
		}
		lastSeen[x] = i + 1
	}
	now := a[len(a) - 1]
	id := len(a)
	next := 0

	for id < 30000000 {
		if lastSeen[now] == 0 {
			next = 0
		} else {
			next = id - lastSeen[now]
		}
		lastSeen[now] = id
		now = next
		id ++
		if id == 2020 {
			ans1 = now
		} else if id == 30000000  {
			ans2 = now
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func processGroup(s string, peopleCount int)(int, int) {
	all := make(map [rune] int)
	for _,x := range s {
		all[x] = all[x] + 1
	}
	allCorrect := 0
	for _,v := range all {
		if v == peopleCount {
			allCorrect++
		}
	}
	return len(all), allCorrect
}

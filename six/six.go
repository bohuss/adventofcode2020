package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input, err := os.Open("six/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 0
	group := ""
	peopleCount := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)

		if len(sLine) == 0 {
			questionsCount, allCorrect := processGroup(group, peopleCount)
			ans1 += questionsCount
			ans2 += allCorrect
			group = ""
			peopleCount = 0
		} else {
			group += sLine
			peopleCount++
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("twentytwo/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 0
	player := 0
	cards := make([][]int, 2)
	cards[0] = make([]int, 0)
	cards[1] = make([]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)
		if strings.HasPrefix(sLine, "Player") {
			continue
		}
		if len(sLine) == 0 {
			player++
		} else {
			num, err := strconv.Atoi(sLine)
			if err != nil {
				panic("unable to parse card")
			}
			cards[player] = append(cards[player], num)
		}
	}

	duplicate := make([][]int, len(cards))
	for i := range cards {
		duplicate[i] = make([]int, len(cards[i]))
		copy(duplicate[i], cards[i])
	}

	for len(cards[0]) > 0 && len(cards[1]) > 0 {
		winner := 1
		if cards[0][0] > cards[1][0] {
			winner = 0
		}
		loser := 1 - winner
		cards[winner] = append(cards[winner][1:], cards[winner][0], cards[loser][0])
		cards[loser] = cards[loser][1:]
	}


	for i, v := range cards[0] {
		ans1 += (len(cards[0]) - i) * v
	}

	for i, v := range cards[1] {
		ans1 += (len(cards[1]) - i) * v
	}

	for i := range cards {
		cards[i] = make([]int, len(duplicate[i]))
		copy(cards[i], duplicate[i])
	}

	for len(cards[0]) > 0 && len(cards[1]) > 0 {
		winner := 1
		if cards[0][0] < len(cards[0]) && cards[1][0] < len(cards[1]) {
			c1 := make([]int, cards[0][0])
			copy(c1, cards[0][1:])
			c2 := make([]int, cards[1][0])
			copy(c2, cards[1][1:])
			winner = playRecursive(c1, c2)
		} else if cards[0][0] > cards[1][0] {
			winner = 0
		}
		loser := 1 - winner
		cards[winner] = append(cards[winner][1:], cards[winner][0], cards[loser][0])
		cards[loser] = cards[loser][1:]
	}

	for i, v := range cards[0] {
		ans2 += (len(cards[0]) - i) * v
	}

	for i, v := range cards[1] {
		ans2 += (len(cards[1]) - i) * v
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func playRecursive(c1 []int, c2 []int) int {
	seen := make(map[string] bool)
	for len(c1) > 0 && len(c2) > 0 {
		code := fmt.Sprintf("%s#%s", fmt.Sprint(c1), fmt.Sprint(c2))
		if seen[code] {
			return 0
		}
		seen[code] = true
		winner := 1
		if c1[0] < len(c1) && c2[0] < len(c2) {
			cc1 := make([]int, c1[0])
			copy(cc1, c1[1:])
			cc2 := make([]int, c2[0])
			copy(cc2, c2[1:])
			winner = playRecursive(cc1[1:], cc2[1:])
		} else if c1[0] > c2[0] {
			winner = 0
		}
		if winner == 0 {
			c1 = append(c1[1:], c1[0], c2[0])
			c2 = c2[1:]
		} else {
			c2 = append(c2[1:], c2[0], c1[0])
			c1 = c1[1:]
		}
	}
	if len(c1) > 0 {
		return 0
	}
	return 1
}

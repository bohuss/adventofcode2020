package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {

	dx := []int {0, 1, 1, 0, -1, -1}
	dy := []int {-1, -1, 0, 1, 1, 0}

	input, err := os.Open("twentyfour/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	ans1 := 0
	ans2 := 0

	p := make([][]bool, 300)
	for i := range p {
		p[i] = make([]bool, 300)
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)
		x, y := 0, 0
		for i, c := range sLine {
			switch c {
			case'n':
				y--
				if sLine[i + 1] == 'w' {
					x++
				}
			case's':
				if sLine[i + 1] == 'e' {
					x--
				}
				y++
			case'e':
				x++
			case'w':
				x--
			}
		}
		p[y + 150][x + 150] = !p[y + 150][x + 150]

	}
	for i := range p {
		for j := range p[0] {
			if p[i][j] {
				ans1++
			}
		}
	}

	p2 := make([][]bool, 300)
	for i := range p {
		p2[i] = make([]bool, 300)
	}

	for day :=0; day <100; day++ {
		for i := range p {
			for j := range p[0] {
				ngh := 0
				for d := range dx {
					y := i + dy[d]
					x := j + dx[d]
					if y<0 || y >= len(p) || x < 0 || x >= len(p[0]){
						continue
					}
					if p[y][x] {
						ngh++
					}
				}
				p2[i][j] = ngh == 2 || (ngh == 1 && p[i][j])

			}
		}
		tmp := p
		p = p2
		p2 = tmp
	}

	for i := range p {
		for j := range p[0] {
			if p[i][j] {
				ans2++
			}
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}


package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {

	input, err := os.Open("eleven/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}

	dx := []int{-1, 0, 1, 0, -1, -1, 1,  1}
	dy := []int{0, -1, 0, 1, -1,  1, 1, -1}

	ans1 := 0
	ans2 := 0
	p := make([][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)
		p = append(p, make([]int, len(sLine)))
		for _, c := range sLine {
			now := 0
			if c == 'L' {
				now = -1
			}
			p[len(p) - 1] = append(p[len(p)-1], now)
		}
	}

	R := len(p)
	C := len(p[0])

	pOrig := make([][]int, R)

	for i := range p {
		pOrig[i] = make([]int, C)
		for j := range p[0] {
			pOrig[i][j] = p[i][j]
		}
	}

	change := true
	for change {
		change = false
		p2 := make([][]int, R)
		for id := range p2 {
			p2[id] = make([]int, C)
		}
		for i := range p2 {
			for j:= range p2[0] {
				c := 0
				for d := range dx {
					x := j + dx[d]
					y := i + dy[d]
					if x < 0 || x == C || y < 0 || y == R {
						continue
					}
					if p[y][x] == 1 {
						c++
					}
				}
				p2[i][j] = p[i][j]
				if p[i][j] == -1 && c == 0 {
					change = true
					p2[i][j] = 1
				} else if p[i][j] == 1 && c > 3 {
					change = true
					p2[i][j] = -1
				}
			}
		}
		p = p2
	}



	for i := range p {
		for j := range p[0] {
			if p[i][j] == 1 {
				ans1++
			}
			p[i][j] = pOrig[i][j]
		}
	}
	ngh := make( [][][]int, R)

	for i := range p {
		ngh[i] = make( [][]int, C)
		for j := range p[0] {
			ngh[i][j] = make([]int, len(dx))
			for d := range dx {
				L := 1
				for x, y := j + dx[d] * L, i + dy[d] * L; y >= 0 && y < R && x >= 0 && x < C && p[y][x] == 0; {
					L++
					x, y = j + dx[d] * L, i + dy[d] * L
				}
				if j + dx[d] * L < 0 || j + dx[d] * L >= C || i + dy[d] * L < 0 || i + dy[d] * L >= R{
					L = -1
				}
				ngh[i][j][d] = L
			}
		}
	}

	change = true
	for change {
		change = false
		p2 := make([][]int, R)
		for id := range p2 {
			p2[id] = make([]int, C)
		}
		for i := range p2 {
			for j:= range p2[0] {
				c := 0
				for d := range dx {
					if ngh[i][j][d] == -1 {
						continue
					}
					x := j + dx[d] * ngh[i][j][d]
					y := i + dy[d] * ngh[i][j][d]

					if p[y][x] == 1 {
						c++
					}
				}
				p2[i][j] = p[i][j]
				if p[i][j] == -1 && c == 0 {
					change = true
					p2[i][j] = 1
				} else if p[i][j] == 1 && c > 4 {
					change = true
					p2[i][j] = -1
				}
			}
		}
		p = p2
	}

	for i := range p {
		for j := range p[0] {
			if p[i][j] == 1 {
				ans2++
			}
		}
	}
	
	fmt.Println(ans1)
	fmt.Println(ans2)
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input, err := os.Open("seventeen/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 0
	ans2 := 0

	c := make([][][]bool, 30)
	c2 := make([][][]bool, 30)
	for i := range c {
		c[i] = make([][]bool, 30)
		c2[i] = make([][]bool, 30)
		for j := range c[0] {
			c[i][j] = make([]bool, 30)
			c2[i][j] = make([]bool, 30)
		}
	}

	h := make([][][][]bool, 30)
	h2 := make([][][][]bool, 30)
	for i := range h {
		h[i] = make([][][]bool, 30)
		h2[i] = make([][][]bool, 30)
		for j := range h[0] {
			h[i][j] = make([][]bool, 30)
			h2[i][j] = make([][]bool, 30)
			for k := range h[0][0] {
				h[i][j][k] = make([]bool, 30)
				h2[i][j][k] = make([]bool, 30)
			}
		}
	}

	for i := 0; i<8; i++{
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		sLine := string(line)
		for j := 0; j <8; j++ {
			c[15][11+i][11+j] = sLine[j] == '#'
			h[15][15][11+i][11+j] = sLine[j] == '#'
		}
	}

	for t := 0; t < 6; t++ {
		for i := range c {
			for j := range c[0] {
				for k := range c[0][0] {
					n := 0
					for ii := -1; ii <= 1; ii ++ {
						for jj := -1; jj <= 1; jj ++ {
							for kk := -1; kk <= 1; kk ++ {
								if ii == 0 && jj == 0 && kk == 0 {
									continue
								}
								if i + ii < 0 || i + ii >= len(c) ||
									j + jj < 0 || j + jj >= len(c[0]) ||
									k + kk < 0 || k + kk >= len(c[0][0]){
									continue
								}
								if c[i + ii][j + jj][k + kk] {
									n++
								}
							}
						}
					}
					c2[i][j][k] =  n == 3 || (c[i][j][k] && n == 2)

					for l := range h[0][0][0] {

						n := 0
						for ii := -1; ii <= 1; ii ++ {
							for jj := -1; jj <= 1; jj ++ {
								for kk := -1; kk <= 1; kk ++ {
									for ll := -1; ll <= 1; ll ++ {
										if ii == 0 && jj == 0 && kk == 0 && ll == 0{
											continue
										}
										if i+ii < 0 || i+ii >= len(c) ||
											j+jj < 0 || j+jj >= len(c[0]) ||
											k+kk < 0 || k+kk >= len(c[0][0]) ||
											l + ll < 0 || l + ll >= len(h[0][0][0]){
											continue
										}
										if h[i+ii][j+jj][k+kk][l+ll] {
											n++
										}
									}
								}
							}
						}
						h2[i][j][k][l] =  n == 3 || (h[i][j][k][l] && n == 2)
					}
				}
			}
		}
		tmp := c2
		c2 = c
		c = tmp

		tmp2 := h2
		h2 = h
		h = tmp2

	}

	for i:= range c {
		for j := range c[0] {
			for k := range c[0][0] {
				if c[i][j][k] {
					ans1++
				}
				for l := range h[0][0][0]{
					if h[i][j][k][l] {
						ans2++
					}
				}
			}
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

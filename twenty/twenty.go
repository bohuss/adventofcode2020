package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Tile struct {
	id int
	lines [] string
	sides [] string
}

type Point struct {
	x int
	y int
}

func (t *Tile) rotate() {
	n := len(t.lines)
	newLines := make([]string, n)
	for i := range newLines {
		newLines[i] = ""
		for j := range t.lines {
			newLines[i] = fmt.Sprintf("%s%c", newLines[i], t.lines[len(t.lines) - j - 1][i])
		}
	}
	t.lines = newLines
	t.sides = getSides(t.lines)
}

func (t *Tile) flip() {
	n := len(t.lines)
	newLines := make([]string, n)
	for i := range newLines {
		newLines[i] = ""
		for j := range t.lines {
			newLines[i] = fmt.Sprintf("%s%c", newLines[i], t.lines[i][len(t.lines) - j - 1])
		}
	}
	t.lines = newLines
	t.sides = getSides(t.lines)
}

func (t *Tile) position(id int, side string) {
	for i := 0; i <8; i++ {
		if t.sides[id] == side {
			break
		}
		t.rotate()
		if i == 4 {
			t.flip()
		}

	}
}

func main() {

	input, err := os.Open("twenty/input.txt")

	reader := bufio.NewReader(input)

	if err != nil {
		panic(err)
	}


	ans1 := 1
	ans2 := int(1e9)
	lines := make([]string, 0)
	tiles := make([]Tile, 0)
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
			num, err := strconv.Atoi(lines[0][5:9])
			if err != nil {
				panic("unable to parse tile id")
			}
			tile := Tile {
				id: num,
				lines: lines[1:],
				sides: getSides(lines[1:]),
			}
			tiles = append(tiles, tile)
			lines = make([]string, 0)
		} else {
			lines = append(lines, sLine)
		}
	}

	oddLines := make(map[string] bool)

	for _, t := range tiles {
		for _, s := range t.sides {
			now := optimize(s)
			oddLines[now] = !oddLines[now]
		}
	}

	cornerTiles := make([]Tile, 0)

	for _, t := range tiles {
		cnt := 0
		for _, s := range t.sides {
			now := optimize(s)
			if oddLines[now] {
				cnt++
			}
		}
		if cnt == 2 {
			ans1 *= t.id
			cornerTiles = append(cornerTiles, t)
		}
	}

	// try all corner tiles in upper left corner in both flipped states
	for _, firstTile := range cornerTiles {
		for _, flip := range []bool{false, true} {
			if flip {
				firstTile.flip()
			}
			board := make([][]*Tile, 0)
			tilesUsedInLevel2 := make(map[int]bool)

			// construct first Row
			board = append(board, make([]*Tile, 0))

			board[0] = append(board[0], &firstTile)
			tilesUsedInLevel2[firstTile.id] = true
			for oddLines[optimize(firstTile.sides[1])] || oddLines[optimize(firstTile.sides[2])] {
				firstTile.rotate()
			}
			currentSide := firstTile.sides[1]

			for true {
				nextTile, err := getTileWithCommonSide(currentSide, tiles, tilesUsedInLevel2)
				if err != nil {
					panic(err)
				}
				tilesUsedInLevel2[nextTile.id] = true
				nextTile.position(3, currentSide)
				board[0] = append(board[0], &nextTile)

				isCornerTile := false
				for _, t := range cornerTiles {
					if t.id == nextTile.id {
						isCornerTile = true
					}
				}
				if isCornerTile {
					break
				}

				currentSide = nextTile.sides[1]

			}

			//construct all other rows
			for row := 1; row < len(board[0]); row++ {
				board = append(board, make([]*Tile, 0))

				currentTile, err := getTileWithCommonSide(board[row-1][0].sides[2], tiles, tilesUsedInLevel2)
				if err != nil {
					panic(err)
				}
				tilesUsedInLevel2[currentTile.id] = true
				currentTile.position(0, board[row-1][0].sides[2])
				board[row] = append(board[row], &currentTile)

				for col := 1; col < len(board[0]); col++ {
					nextTile, err := getTileWithCommonSide(board[row][col-1].sides[1], tiles, tilesUsedInLevel2)
					if err != nil {
						panic(err)
					}
					tilesUsedInLevel2[nextTile.id] = true

					nextTile.position(3, board[row][col-1].sides[1])

					board[row] = append(board[row], &nextTile)
				}
			}

			// construct painting (glue pieces together without frames)
			width := len(board[0]) * (len(board[0][0].lines[0]) - 2)
			height := width
			canvas := make([][]bool, height)
			for i := range canvas {
				canvas[i] = make([]bool, width)
			}

			boardRow := 0
			boardCol := 0
			row := 1

			for i := range canvas {
				col := 1
				for j := range canvas[0] {
					canvas[i][j] = board[boardRow][boardCol].lines[row][col] == '#'
					col++
					if col == 9 {
						boardCol++
						col = 1
					}
				}
				row++
				boardCol = 0
				if row == 9 {
					boardRow++
					row = 1
				}
			}

			// find monsters in canvas
			monsterPattern := []string{
				"                  # ",
				"#    ##    ##    ###",
				" #  #  #  #  #  #   ",
			}
			monsterCoordinates := make([]Point, 0)
			for i := range monsterPattern {
				for j := range monsterPattern[0] {
					if monsterPattern[i][j] == '#' {
						monsterCoordinates = append(monsterCoordinates, Point{x: j, y: i})
					}
				}
			}

			monstersCount := 0
			allCount := 0
			for i := range canvas {
				for j := range canvas {
					if canvas[i][j] {
						allCount++
					}
					if isMonsterLocated(canvas, i, j, monsterCoordinates) {
						monstersCount += len(monsterCoordinates)
					}
				}
			}

			if allCount-monstersCount < ans2 {
				ans2 = allCount - monstersCount
			}

			tilesUsedInLevel2 = make(map[int]bool)
		}

	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

func isMonsterLocated(canvas [][]bool, i int, j int, coordinates []Point) bool {

	for _, v :=range coordinates {
		if i + v.y >= len(canvas) || j + v.x >= len(canvas[0]) || !canvas[i + v.y][j + v.x] {
			return false
		}
	}
	return true
}

func getTileWithCommonSide(side1 string, tiles []Tile, usedTiles map[int]bool) (Tile, error) {
	for _,tile2 := range tiles {
		if usedTiles[tile2.id] {
			continue
		}
		for _, side2 := range tile2.sides {
			if optimize(side1) == optimize(side2) {
				return  tile2, nil
			}
		}
	}
	return Tile{}, errors.New("no tile found")
}

func optimize(s string) string {
	alt := ""
	for i := range s {
		alt += string(s[len(s) - i - 1])
	}
	if s < alt {
		return s
	}
	return alt
}

func getSides(lines []string) []string {
	linesRight := ""
	linesLeft := ""
	for i := range lines {
		linesRight = fmt.Sprintf("%s%c", linesRight, lines[i][9])
		linesLeft = fmt.Sprintf("%s%c", linesLeft, lines[i][0])
	}
	ans := []string{lines[0], linesRight, lines[len(lines) - 1], linesLeft}
	return ans
}

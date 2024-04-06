package life

import "fmt"

func CreateSeed(patternStr string, xSize, ySize int) [][]uint8 {
	pattern, ok := patterns[patternStr]
	if !ok {
		panic("unknown pattern - " + patternStr)
	}
	seed := make([][]uint8, ySize, ySize)

	var patternXSize, patternYSize int
	patternYSize = len(pattern)
	for _, row := range pattern {
		if len(row) > patternXSize {
			patternXSize = len(row)
		}
	}

	startX := (xSize - patternXSize) / 2
	startY := (ySize - patternYSize) / 2

	for y := range seed {
		seed[y] = make([]uint8, xSize, xSize)
	}
	fmt.Println("pattern", pattern)
	fmt.Println("xSize", xSize, "ySize", ySize)

	fmt.Println("startX", startX, "startY", startY)
	fmt.Println("patternXSize", patternXSize, "patternYSize", patternYSize)

	for x := 0; x <= patternXSize; x++ {
		for y := 0; y <= patternYSize; y++ {
			seed[y+startY][x+startX] = getValue(pattern, y, x)
		}
	}

	return seed
}

func getValue(pattern [][]uint8, x, y int) uint8 {
	if x >= len(pattern) {
		return 0
	}
	if y >= len(pattern[x]) {
		return 0
	}

	return pattern[x][y]

}

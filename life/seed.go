package life

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

	for x := 0; x <= patternXSize; x++ {
		for y := 0; y <= patternYSize; y++ {
			seed[y+startY][x+startX] = getValue(pattern, x, y)
		}
	}

	return seed
}

func getValue(pattern [][]uint8, x, y int) uint8 {
	if y >= len(pattern) {
		return 0
	}
	if x >= len(pattern[y]) {
		return 0
	}

	return pattern[y][x]

}

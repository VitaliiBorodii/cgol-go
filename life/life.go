package life

import (
	"fmt"
	"strings"
	"time"
)

var animationDelay = time.Millisecond * 100

type Life struct {
	seed  [][]uint8
	xSize int
	ySize int

	currentState map[string][2]int

	_calculatedPoints map[string]interface{}
}

func pointKey(x, y int) string {
	return fmt.Sprintf("%dx%d", x, y)
}

func NewLife(seed [][]uint8) *Life {
	l := Life{}
	l.seed = seed
	l.calculateInitState()
	fmt.Println("Grid size: %d x %d\n", l.xSize, l.ySize)
	return &l
}

func (l *Life) Run(cycles int) {
	for i := 0; i <= cycles; i++ {
		fmt.Printf("cycle: %d | alive cells: %d\n", i, len(l.currentState))
		l.printStateString()
		time.Sleep(animationDelay)
		l.nextState()
	}
}

func (l *Life) calculateInitState() {
	state := map[string][2]int{}
	l.ySize = len(l.seed)
	for y, row := range l.seed {
		if len(row) > l.xSize {
			l.xSize = len(row)
		}
		for x, el := range row {
			if el == 1 {
				state[pointKey(x, y)] = [2]int{x, y}
			}
		}
	}
	l.currentState = state
	l._calculatedPoints = map[string]interface{}{}
}

func (l *Life) nextState() {
	newState := map[string][2]int{}

	for _, point := range l.currentState {
		subgrid := l.pointBoundaries(point)

		for x := subgrid[0][0]; x <= subgrid[1][0]; x++ {
			for y := subgrid[0][1]; y <= subgrid[1][1]; y++ {
				newPoint := [2]int{x, y}
				key := pointKey(x, y)

				// no need to check point that was already checked
				if _, ok := l._calculatedPoints[key]; ok {
					continue
				}

				count := l.countAliveNeighbours(newPoint)
				_, alive := l.currentState[key]

				if alive {
					if count == 2 || count == 3 {
						newState[key] = newPoint
					}
				} else if count == 3 {
					newState[key] = newPoint
				}
				l._calculatedPoints[key] = nil
			}
		}
	}
	l.currentState = newState
	l._calculatedPoints = map[string]interface{}{}
}

func (l *Life) countAliveNeighbours(point [2]int) int {
	startKey := pointKey(point[0], point[1])
	subgrid := l.pointBoundaries(point)
	var neighboursCount int

	for x := subgrid[0][0]; x <= subgrid[1][0]; x++ {
		for y := subgrid[0][1]; y <= subgrid[1][1]; y++ {
			key := pointKey(x, y)
			if key == startKey { // do not count itself
				continue
			}
			if _, ok := l.currentState[key]; ok {
				neighboursCount++
			}
		}
	}
	return neighboursCount
}

func (l *Life) pointBoundaries(point [2]int) [2][2]int {
	x, y := point[0], point[1]
	var result [2][2]int

	if x-1 < 0 { // xMin
		result[0][0] = x
	} else {
		result[0][0] = x - 1
	}
	if x+1 >= l.xSize { // xMax
		result[1][0] = x
	} else {
		result[1][0] = x + 1
	}

	if y-1 < 0 { // yMin
		result[0][1] = y
	} else {
		result[0][1] = y - 1
	}
	if y+1 >= l.ySize { // yMax
		result[1][1] = y
	} else {
		result[1][1] = y + 1
	}
	return result
}

func (l *Life) printStateString() {
	str := l.getStateString()
	fmt.Println(str)
}

func (l *Life) getStateString() string {
	canv := strings.Repeat("_", l.xSize*2+2)
	for y := 0; y < l.ySize; y++ {
		canv += "\n|"
		for x := 0; x < l.xSize; x++ {
			key := pointKey(x, y)
			_, ok := l.currentState[key]
			if !ok {
				canv += "  "
			} else {
				canv += " *"
			}
		}
		canv += "|"
	}
	canv += "\n"
	canv += strings.Repeat("-", l.xSize*2+2)
	canv += "\n"
	return canv
}

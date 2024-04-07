package life

import (
	"fmt"
	"strings"
	"time"
)

var animationDelay = time.Second

func pointKey(x, y int) string {
	return fmt.Sprintf("%dx%d", x, y)
}

type Life struct {
	seed  [][]uint8
	xSize int
	ySize int
	cycle int

	currentState map[string][2]int

	withOverflow   bool
	animationDelay time.Duration

	calculatedPoints map[string]interface{}
}

func NewLife(seed [][]uint8) *Life {
	l := Life{}
	l.seed = seed
	l.animationDelay = animationDelay
	l.calculateInitState()
	fmt.Printf("Grid size: %d x %d\n", l.xSize, l.ySize)
	return &l
}

func (l *Life) tick() {
	fmt.Printf("cycle: %d | alive cells: %d\n", l.cycle, len(l.currentState))
	l.printStateString()
	time.Sleep(l.animationDelay)
	l.nextState()
	l.cycle++
}

func (l *Life) Run(cycles int) {
	for i := 0; i <= cycles; i++ {
		l.tick()
	}
}

func (l *Life) RunInfinite() {
	for {
		l.tick()
	}
}

func (l *Life) WithOverFlow(overflow bool) {
	l.withOverflow = overflow
}

func (l *Life) WithAnimationSpeed(speed uint) {
	l.animationDelay = animationDelay / time.Duration(speed)
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
	l.calculatedPoints = map[string]interface{}{}
}

func (l *Life) nextState() {
	newState := map[string][2]int{}

	for _, point := range l.currentState {
		neighbours := l.getNeighrouringPoints(point)

		for _, point := range append(neighbours, point) {
			x, y := point[0], point[1]

			newPoint := [2]int{x, y}
			key := pointKey(x, y)

			// no need to check point that was already checked
			if _, ok := l.calculatedPoints[key]; ok {
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
			l.calculatedPoints[key] = nil
		}
	}
	l.currentState = newState
	l.calculatedPoints = map[string]interface{}{}
}

func (l *Life) countAliveNeighbours(point [2]int) int {
	neighbouringPoints := l.getNeighrouringPoints(point)
	var neighboursCount int

	for _, point := range neighbouringPoints {
		x, y := point[0], point[1]

		key := pointKey(x, y)

		if _, ok := l.currentState[key]; ok {
			neighboursCount++
		}
	}

	return neighboursCount
}

func (l *Life) getNeighrouringPoints(point [2]int) [][2]int {
	x, y := point[0], point[1]
	var result = make([][2]int, 0, 8)
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}

			if !l.withOverflow {
				if i < 0 || i >= l.xSize || j < 0 || j >= l.ySize {
					continue
				}
			}

			var xPoint, yPoint int
			if i < 0 {
				xPoint = i + l.xSize
			} else if i >= l.xSize {
				xPoint = i - l.xSize
			} else {
				xPoint = i
			}

			if j < 0 {
				yPoint = j + l.ySize
			} else if j >= l.ySize {
				yPoint = j - l.ySize
			} else {
				yPoint = j
			}

			result = append(result, [2]int{xPoint, yPoint})
		}
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

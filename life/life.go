package life

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type life struct {
	gridWidth  int32
	gridHeight int32

	grid    [][]bool
	newGrid [][]bool

	generation int
	liveCells  int

	screenWidth  int32
	screenHeight int32

	cellSize int32
}

func NewLife(screenWidth, screenHeight, cellSize int32) *life {
	gridW := screenWidth / cellSize
	gridH := screenHeight / cellSize

	grid := make([][]bool, gridW)
	newGrid := make([][]bool, gridW)

	for i := range grid {
		grid[i] = make([]bool, gridH)
		newGrid[i] = make([]bool, gridH)
	}

	return &life{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		gridWidth:    gridW,
		gridHeight:   gridH,
		grid:         grid,
		newGrid:      newGrid,
		cellSize:     cellSize,
		generation:   0,
		liveCells:    0,
	}
}

func (l *life) Start() {
	defer rl.CloseWindow()

	rl.InitWindow(l.screenWidth, l.screenHeight, "Игра в жизнь")

	l.initGrid()
	rl.SetTargetFPS(10)

	backgroundColor := rl.NewColor(30, 30, 30, 255)

	l.generation = 0

	for !rl.WindowShouldClose() {
		l.updateGrid()
		l.generation++

		l.liveCells = l.countLiveCells()

		windowTitle := fmt.Sprintf("Игра в жизнь - Генерация: %d | Живых клеток: %d", l.generation, l.liveCells)
		rl.SetWindowTitle(windowTitle)

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		l.drawGrid()

		rl.EndDrawing()
	}
}

func (l *life) initGrid() {
	for i := range l.gridWidth {
		for j := range l.gridHeight {
			l.grid[i][j] = rl.GetRandomValue(0, 1) == 1
		}
	}
}

func (l *life) updateGrid() {
	for i := range l.gridWidth {
		for j := range l.gridHeight {
			l.newGrid[i][j] = l.grid[i][j]
		}
	}

	for i := range l.gridWidth {
		for j := range l.gridHeight {
			liveNeighbors := l.countLiveNeighbors(i, j)

			if l.grid[i][j] {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					l.newGrid[i][j] = false
				}
			} else {
				if liveNeighbors == 3 {
					l.newGrid[i][j] = true
				}
			}
		}
	}

	for i := range l.gridWidth {
		for j := range l.gridHeight {
			l.grid[i][j] = l.newGrid[i][j]
		}
	}
}

func (l *life) countLiveNeighbors(x, y int32) int32 {
	var count int32 = 0
	var i, j int32
	for i = -1; i <= 1; i++ {
		for j = -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			neighborX := x + i
			neighborY := y + j

			if neighborX >= 0 && neighborX < l.gridWidth && neighborY >= 0 && neighborY < l.gridHeight {
				if l.grid[neighborX][neighborY] {
					count++
				}
			}
		}
	}
	return count
}

func (l *life) countLiveCells() int {
	count := 0
	for i := range l.gridWidth {
		for j := range l.gridHeight {
			if l.grid[i][j] {
				count++
			}
		}
	}
	return count
}

func (l *life) drawGrid() {
	for i := range l.gridWidth {
		for j := range l.gridHeight {
			if l.grid[i][j] {
				rl.DrawRectangle(int32(i*l.cellSize), int32(j*l.cellSize), int32(l.cellSize), int32(l.cellSize), rl.DarkGreen)
			}
		}
	}
}

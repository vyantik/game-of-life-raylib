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

	originalScreenWidth  int32
	originalScreenHeight int32

	cellSize int32

	previousLiveCells int
	stableCycles      int

	isFullscreen bool
	isGameOver   bool

	previousGrids [][][]bool
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
		screenWidth:          screenWidth,
		screenHeight:         screenHeight,
		originalScreenWidth:  screenWidth,
		originalScreenHeight: screenHeight,
		gridWidth:            gridW,
		gridHeight:           gridH,
		grid:                 grid,
		newGrid:              newGrid,
		cellSize:             cellSize,
		generation:           0,
		liveCells:            0,
		isFullscreen:         false,
		isGameOver:           false,
		previousLiveCells:    0,
		stableCycles:         0,
	}
}

func (l *life) Start() {
	defer rl.CloseWindow()

	rl.InitWindow(l.screenWidth, l.screenHeight, "Игра в жизнь")

	l.initGrid()
	rl.SetTargetFPS(0)

	backgroundColor := rl.NewColor(30, 30, 30, 255)

	for !rl.WindowShouldClose() {
		l.handleControls()

		l.handleGameOver()

		windowTitle := fmt.Sprintf("Игра в жизнь - Генерация: %d | Живых клеток: %d", l.generation, l.liveCells)
		rl.SetWindowTitle(windowTitle)

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		l.drawGrid()

		if l.isFullscreen {
			l.drawUI()
		} else {
			l.drawHint()
		}

		if l.isGameOver {
			l.drawGameOver()
		}

		rl.EndDrawing()
	}
}

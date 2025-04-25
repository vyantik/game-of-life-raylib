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

	isFullscreen bool
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
	}
}

func (l *life) Start() {
	defer rl.CloseWindow()

	rl.InitWindow(l.screenWidth, l.screenHeight, "Игра в жизнь")

	l.initGrid()
	rl.SetTargetFPS(60)

	backgroundColor := rl.NewColor(30, 30, 30, 255)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyF) {
			l.toggleFullscreen()
		}

		l.updateGrid()
		l.generation++

		l.liveCells = l.countLiveCells()

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

		rl.EndDrawing()
	}
}

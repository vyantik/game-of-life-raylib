package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 1000
	cellSize     = 10
	gridWidth    = screenWidth / cellSize
	gridHeight   = screenHeight / cellSize
)

var (
	grid    [gridWidth][gridHeight]bool
	newGrid [gridWidth][gridHeight]bool

	generation int
	liveCells  int
)

func main() {
	defer rl.CloseWindow()

	rl.InitWindow(screenWidth, screenHeight, "Игра в жизнь")

	initGrid()
	rl.SetTargetFPS(10)

	backgroundColor := rl.NewColor(30, 30, 30, 255)

	generation = 0

	for !rl.WindowShouldClose() {
		updateGrid()
		generation++

		liveCells = countLiveCells()

		windowTitle := fmt.Sprintf("Игра в жизнь - Генерация: %d | Живых клеток: %d", generation, liveCells)
		rl.SetWindowTitle(windowTitle)

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		drawGrid()

		rl.EndDrawing()
	}
}

func initGrid() {
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			grid[i][j] = rl.GetRandomValue(0, 1) == 1
		}
	}
}

func updateGrid() {
	for i := range gridWidth {
		for j := 0; j < gridHeight; j++ {
			newGrid[i][j] = grid[i][j]
		}
	}

	for i := range gridWidth {
		for j := 0; j < gridHeight; j++ {
			liveNeighbors := countLiveNeighbors(i, j)

			if grid[i][j] {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					newGrid[i][j] = false
				}
			} else {
				if liveNeighbors == 3 {
					newGrid[i][j] = true
				}
			}
		}
	}

	for i := range gridWidth {
		for j := range gridHeight {
			grid[i][j] = newGrid[i][j]
		}
	}
}

func countLiveNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			neighborX := x + i
			neighborY := y + j

			if neighborX >= 0 && neighborX < gridWidth && neighborY >= 0 && neighborY < gridHeight {
				if grid[neighborX][neighborY] {
					count++
				}
			}
		}
	}
	return count
}

func countLiveCells() int {
	count := 0
	for i := range gridWidth {
		for j := range gridHeight {
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}

func drawGrid() {
	for i := range gridWidth {
		for j := range gridHeight {
			if grid[i][j] {
				rl.DrawRectangle(int32(i*cellSize), int32(j*cellSize), int32(cellSize), int32(cellSize), rl.DarkGreen)
			}
		}
	}
}

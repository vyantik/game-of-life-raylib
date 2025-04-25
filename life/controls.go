package life

import rl "github.com/gen2brain/raylib-go/raylib"

func (l *life) toggleFullscreen() {
	if !l.isFullscreen {
		rl.SetWindowSize(1920, 1080)
		rl.ToggleFullscreen()
		l.isFullscreen = true

		currentScreenWidth := int32(rl.GetScreenWidth())
		currentScreenHeight := int32(rl.GetScreenHeight())

		l.screenWidth = currentScreenWidth
		l.screenHeight = currentScreenHeight

		l.gridWidth = l.screenWidth / l.cellSize
		l.gridHeight = l.screenHeight / l.cellSize

		l.resizeGrid()
		l.generation = 0
		l.liveCells = l.countLiveCells()
	} else {
		rl.ToggleFullscreen()
		rl.SetWindowSize(int(l.originalScreenWidth), int(l.originalScreenHeight))
		l.isFullscreen = false

		currentScreenWidth := int32(rl.GetScreenWidth())
		currentScreenHeight := int32(rl.GetScreenHeight())

		l.screenWidth = currentScreenWidth
		l.screenHeight = currentScreenHeight

		l.gridWidth = l.screenWidth / l.cellSize
		l.gridHeight = l.screenHeight / l.cellSize

		l.resizeGrid()
		l.generation = 0
		l.liveCells = l.countLiveCells()

	}
}

func (l *life) handleControls() {
	if rl.IsKeyPressed(rl.KeyF) {
		l.resetGame()
		l.toggleFullscreen()
	}

	if rl.IsKeyPressed(rl.KeyR) && l.isGameOver {
		l.resetGame()
	}
}

func (l *life) resetGame() {
	l.generation = 0
	l.liveCells = 0
	l.isGameOver = false
	l.previousLiveCells = 0
	l.stableCycles = 0
	l.initGrid()
}

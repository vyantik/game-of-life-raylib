package life

import rl "github.com/gen2brain/raylib-go/raylib"

func (l *life) resizeGrid() {
	l.grid = make([][]bool, l.gridWidth)
	l.newGrid = make([][]bool, l.gridWidth)
	for i := range l.grid {
		l.grid[i] = make([]bool, l.gridHeight)
		l.newGrid[i] = make([]bool, l.gridHeight)
	}
	l.initGrid()
}

func (l *life) initGrid() {
	if len(l.grid) != int(l.gridWidth) || (l.gridWidth > 0 && len(l.grid[0]) != int(l.gridHeight)) {
		l.grid = make([][]bool, l.gridWidth)
		for i := range l.grid {
			l.grid[i] = make([]bool, l.gridHeight)
		}
	}

	for i := range l.gridWidth {
		for j := range l.gridHeight {
			l.grid[i][j] = rl.GetRandomValue(0, 1) == 1
		}
	}
}

func (l *life) updateGrid() {
	if len(l.newGrid) != int(l.gridWidth) || (l.gridWidth > 0 && len(l.newGrid[0]) != int(l.gridHeight)) {
		l.newGrid = make([][]bool, l.gridWidth)
		for i := range l.newGrid {
			l.newGrid[i] = make([]bool, l.gridHeight)
		}
	}

	if len(l.grid) != int(l.gridWidth) || (l.gridWidth > 0 && len(l.grid[0]) != int(l.gridHeight)) {
		l.grid = make([][]bool, l.gridWidth)
		for i := range l.grid {
			l.grid[i] = make([]bool, l.gridHeight)
		}
	}

	for i := range l.gridWidth {
		for j := range l.gridHeight {
			if int32(i) < l.gridWidth && int32(j) < l.gridHeight {
				l.newGrid[i][j] = l.grid[i][j]
			}
		}
	}

	for i := range l.gridWidth {
		for j := range l.gridHeight {
			if int32(i) < l.gridWidth && int32(j) < l.gridHeight {
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
	}

	if len(l.grid) != int(l.gridWidth) || (l.gridWidth > 0 && len(l.grid[0]) != int(l.gridHeight)) {
		l.grid = make([][]bool, l.gridWidth)
		for i := range l.grid {
			l.grid[i] = make([]bool, l.gridHeight)
		}
	}

	for i := range l.gridWidth {
		for j := range l.gridHeight {
			if int32(i) < l.gridWidth && int32(j) < l.gridHeight {
				l.grid[i][j] = l.newGrid[i][j]
			}
		}
	}
}

func (l *life) drawGrid() {
	for i := range l.gridWidth {
		for j := range l.gridHeight {
			if int32(i) < l.gridWidth && int32(j) < l.gridHeight {
				if l.grid[i][j] {
					rl.DrawRectangle(int32(i*l.cellSize), int32(j*l.cellSize), int32(l.cellSize), int32(l.cellSize), rl.DarkGreen)
				}
			}
		}
	}
}

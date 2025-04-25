package life

func (l *life) handleGameOver() {
	if !l.isGameOver {
		prevGrid := make([][]bool, l.gridWidth)
		for i := range prevGrid {
			prevGrid[i] = make([]bool, l.gridHeight)
			copy(prevGrid[i], l.grid[i])
		}

		l.updateGrid()
		l.generation++
		l.liveCells = l.countLiveCells()

		if l.liveCells == 0 {
			l.isGameOver = true
		} else if l.liveCells == l.previousLiveCells {
			diffCells := 0
			totalCells := 0

			for i := int32(0); i < l.gridWidth; i++ {
				for j := int32(0); j < l.gridHeight; j++ {
					if i < int32(len(prevGrid)) && j < int32(len(prevGrid[i])) &&
						i < int32(len(l.grid)) && j < int32(len(l.grid[i])) {
						totalCells++
						if prevGrid[i][j] != l.grid[i][j] {
							diffCells++
						}
					}
				}
			}

			threshold := totalCells * 10 / 100

			if diffCells <= threshold {
				l.stableCycles++
				if l.stableCycles >= 3 {
					l.isGameOver = true
				}
			} else {
				l.stableCycles = 0
			}
		} else {
			l.stableCycles = 0
		}

		l.previousLiveCells = l.liveCells
	}
}

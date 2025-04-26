package life

import "reflect"

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
			return
		}

		for _, previousState := range l.previousGrids {
			if reflect.DeepEqual(previousState, l.grid) {
				l.isGameOver = true
				return
			}
		}

		if len(l.previousGrids) >= 10 {
			l.previousGrids = l.previousGrids[1:]
		}
		l.previousGrids = append(l.previousGrids, prevGrid)
		if l.previousLiveCells == l.liveCells {
			l.stableCycles++
			if l.stableCycles >= 10 {
				l.isGameOver = true
			}
		} else {
			l.stableCycles = 0
		}

		l.previousLiveCells = l.liveCells
	}
}

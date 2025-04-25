package life

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
				if len(l.grid) > int(neighborX) && len(l.grid[neighborX]) > int(neighborY) {
					if l.grid[neighborX][neighborY] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (l *life) countLiveCells() int {
	count := 0
	if len(l.grid) == 0 || (l.gridWidth > 0 && len(l.grid[0]) == 0) {
		return 0
	}
	for i := range l.grid {
		for j := range l.grid[i] {
			if int32(i) < l.gridWidth && int32(j) < l.gridHeight {
				if l.grid[i][j] {
					count++
				}
			}
		}
	}
	return count
}

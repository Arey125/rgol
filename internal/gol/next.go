package gol

func isAlive(state State, i int, j int) bool {
	h := len(state)
	w := len(state[0])

	if i < 0 || i >= h {
		return false
	}

	if j < 0 || j >= w {
		return false
	}

	return state[i][j]
}

func aliveInNextState(state State, i int, j int) bool {
	count := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
            if di == 0 && dj == 0 {
                continue
            }
			if isAlive(state, i+di, j+dj) {
				count++
			}
		}
	}
	if state[i][j] {
		return count == 2 || count == 3
	}
    return count == 3
}

func Next(state State) State {
	h := len(state)
	w := len(state[0])

	next := make([][]bool, h)
	for i := range next {
		next[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			next[i][j] = aliveInNextState(state, i, j)
		}
	}
	return next
}

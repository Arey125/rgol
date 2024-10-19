package gol

func PreviousForOneCell(cell bool) []State {
    res := make([]State, 0, 100)
    for mask := 0; mask < 1 << 9; mask++ {
        state := NewState(3, 3)
        for i := 0; i < 9;i++ {
            state[i/3][i % 3] = (mask & (1 << i) != 0)
        }
        next := Next(state)
        if next[1][1] == cell {
            res = append(res, state)
        }
    }
    return res
}

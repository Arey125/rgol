package gol

import (
    "fmt"
    "strings"
)

type State [][]bool

func NewState(w int, h int) State {
    res := make([][]bool, h)
    for i := range res {
        res[i] = make([]bool, w)
    }
    return res
}

func getWidthAndHeight(state State) (int, int) {
	h := len(state)
	w := len(state[0])
    return w, h
}

func Print(state State) {
    for _, row := range state {
        for _, cell := range row {
            if cell {
                fmt.Print("#");
                continue
            }
            fmt.Print(".")
        }
        fmt.Print("\n")
    }
}

func FromString(stateString string) State {
    lines := strings.Split(stateString, "\n")
    res := NewState(len(lines[0]), len(lines))

    for i, line := range lines {
        for j, c := range line {
            if c == '#' {
                res[i][j] = true
                continue
            }
            res[i][j] = false
        }
    }
    return res
}

package gol

import (
    "fmt"
    "strings"
)

type State [][]bool

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
    res := make([][]bool, len(lines))

    for i, line := range lines {
        res[i] = make([]bool, len(line))
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

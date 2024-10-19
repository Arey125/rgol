package main

import (
	"fmt"

	"github.com/Arey125/rgol/internal/gol"
)

func main() {
    states := gol.PreviousForOneCell(false)
    fmt.Println(len(states))

    for _, state := range states {
        gol.Print(state)
        fmt.Println()
    }
}

package main

import (
    "github.com/Arey125/rgol/internal/gol"
)

const stateString = `...
###
...`

func main() {
    state := gol.FromString(stateString)
    state = gol.Next(state)
    gol.Print(state)
}

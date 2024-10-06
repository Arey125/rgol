package main

import (
	"fmt"

	"github.com/Arey125/rgol/internal/gol"
)

const stateString = `...
###
...`

const anotherStateString = `###
...
###`

func main() {
    state := gol.FromString(stateString)
    state = gol.Next(state)
    anotherState := gol.FromString(anotherStateString)
    fmt.Println(gol.DoMatchVertically(state, anotherState))
}

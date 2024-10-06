package main

import (
	"fmt"

	"github.com/Arey125/rgol/internal/gol"
)

const stateString = `...
###
...`

func main() {
    state := gol.FromString(stateString)
    fmt.Println(gol.DoMatchHorizontally(state, state))
}

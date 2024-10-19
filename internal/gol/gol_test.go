package gol

import (
    "testing"
)

const stateString = `...
###
...`

const anotherStateString = `###
...
###`

func TestMatch(t *testing.T) {
    state := FromString(stateString)
    anotherState := FromString(anotherStateString)
    if !DoMatchVertically(state, anotherState) {
        t.Errorf("states should match")
    }

    state = Next(state)
    if DoMatchVertically(state, anotherState) {
        t.Errorf("states should not match")
    }
}

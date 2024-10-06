package gol

func DoMatchHorizontally(s1 State, s2 State) bool {
    w1, h1 := getWidthAndHeight(s1)
    _, h2 := getWidthAndHeight(s2)
    if (h1 != h2) {
        return false
    }
    for i := 0;i < h1;i++ {
        if s1[i][w1 - 2] != s2[i][0] || s1[i][w1 - 1] != s2[i][1] {
            return false
        }
    }
    return true
}

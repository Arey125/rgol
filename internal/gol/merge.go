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

func DoMatchVertically(s1 State, s2 State) bool {
    w1, h1 := getWidthAndHeight(s1)
    w2, _ := getWidthAndHeight(s2)
    if (w1 != w2) {
        return false
    }
    for i := 0;i < w1;i++ {
        if s1[h1 - 2][i] != s2[0][i] || s1[h1 - 1][i] != s2[1][i] {
            return false
        }
    }
    return true
}

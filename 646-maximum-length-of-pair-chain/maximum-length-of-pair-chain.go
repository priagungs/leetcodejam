func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][0] < pairs[j][0]
    })

    dp := make(map[int]int, len(pairs))

    for i := 0; i < len(pairs); i++ {
        solve(pairs, i, dp)
    }

    res := 0
    for _, val := range dp {
        res = max(res, val)
    }
    return res
}

func solve(pairs [][]int, curr int, dp map[int]int) {
    if _, ok := dp[curr]; ok {
        return
    }

    res := 0
    for i := curr+1; i < len(pairs); i++ {
        if pairs[i][0] > pairs[curr][1] {
            solve(pairs, i, dp)
            res = max(res, dp[i])
        }
    }
    dp[curr] = 1 + res 
}

func max(a, b int) int {
    if a > b {return a}
    return b
}
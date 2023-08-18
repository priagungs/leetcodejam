func maximalNetworkRank(n int, roads [][]int) int {

    graph := make([][]int, n)
    count := make([]int, n)
    for i := 0; i < n; i++ {
        graph[i] = make([]int, n)
    }

    for _, road := range roads {
        graph[road[0]][road[1]] = 1
        graph[road[1]][road[0]] = 1
        count[road[0]]++
        count[road[1]]++
    }

    res := 0
    for i := 0; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            res = max(res, count[i] + count[j] - graph[i][j])
        }
    }

    return res

}

func max(a, b int) int {
    if a > b {return a}
    return b
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                continue
            }

            if i == 0 && j == 0 {
                dp[i][j] = 1
                continue
            }

            if i > 0 {
                dp[i][j] += dp[i-1][j]
            }

            if j > 0 {
                dp[i][j] += dp[i][j-1]
            }
        }
    }

    return dp[m-1][n-1]
}
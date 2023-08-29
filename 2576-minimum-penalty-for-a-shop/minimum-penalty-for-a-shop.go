func bestClosingTime(customers string) int {
    dp := make([]int, len(customers)+1)
    totalCustomers := 0
    
    for i := 0; i < len(customers); i++ {
        if customers[i] == 'Y' {totalCustomers++}
    }

    dp[0] = totalCustomers
    for i := 1; i < len(customers)+1; i++ {
        if customers[i-1] == 'Y' {
            dp[i] = dp[i-1] - 1
        } else {
            dp[i] = dp[i-1] + 1
        }
    }

    bestTime, minPenalty := 0, dp[0]
    for i := 1; i < len(customers)+1; i++ {
        if dp[i] < minPenalty {
            bestTime, minPenalty = i, dp[i]
        }
    }

    return bestTime
}
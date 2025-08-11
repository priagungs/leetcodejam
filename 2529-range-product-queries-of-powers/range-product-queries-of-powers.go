const mod = 1_000_000_007
func productQueries(n int, queries [][]int) []int {
    maxPow := 1
    i := 0
    for maxPow <= n/2 {
        maxPow = maxPow*2    
        i++
    }

    val := n
    powers := make([]int, 0)
    for val > 0 {
        if maxPow > val {
            maxPow = maxPow/2
            i--
            continue
        }
        val -= maxPow
        powers = append([]int{i}, powers...)
    }

    prefSum := make([]int, len(powers))
    prefSum[0] = powers[0]
    for i := 1; i < len(powers); i++ {
        prefSum[i] = prefSum[i-1] + powers[i]
    }

    res := make([]int, 0, len(queries))
    for _, query := range queries {
        valRes := math.Pow(2, float64(prefSum[query[1]]))
        if query[0] > 0 {
            valRes = valRes / math.Pow(2, float64(prefSum[query[0]-1]))
        }
        res = append(res, int(math.Mod(valRes, mod)))
    }

    return res
}
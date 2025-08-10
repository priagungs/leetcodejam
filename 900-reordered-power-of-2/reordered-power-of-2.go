func reorderedPowerOf2(n int) bool {
    numCount := toNumCountTable(n)
    digitNum := countDigit(n)

    listPow2 := make([]int, 0)
    for i := 1; countDigit(i) <= digitNum; i = i*2 {
        if countDigit(i) == digitNum {
            val := i
            listPow2 = append(listPow2, val)
        }
    }

    for _, pow2 := range listPow2 {
        pow2CountTable := toNumCountTable(pow2)
        if evaluate(pow2CountTable, numCount) {
            return true
        }
    }
    return false
}

func evaluate(a, b map[int]int) bool {
    if len(a) != len(b) { return false }
    for k, v := range a {
        if v != b[k] {
            return false
        }
    }
    return true
}

func toNumCountTable(n int) map[int]int {
    res := make(map[int]int)
    for i := n; i > 0; i = i/10 {
        res[i%10]++
    }
    return res
}

func countDigit(n int) int {
    res := 0
    for i := n; i != 0; i = i/10 {
        res++
    }
    return res
}
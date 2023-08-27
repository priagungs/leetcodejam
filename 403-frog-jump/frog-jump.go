type key struct {
    i, k int
}

func canCross(stones []int) bool {
    return backtrack(stones, 0, 0, make(map[key]bool))
}

func backtrack(stones []int, i, k int, mem map[key]bool) bool {
    if val, found := mem[key{i: i, k: k}]; found {
        return val
    }

    if i == len(stones)-1 {
        mem[key{i: i, k: k}] = true
        return true
    }

    if stones[i+1] - stones[i] > k+1 {
        mem[key{i: i, k: k}] = false
        return false
    }

    res := false
    for j := i+1; j < len(stones) && stones[j]-stones[i] <= k+1; j++ {
        if stones[j]-stones[i] == k-1 {
            res = res || backtrack(stones, j, k-1, mem)
        }

        if stones[j]-stones[i] == k {
            res = res || backtrack(stones, j, k, mem)
        }

        if stones[j]-stones[i] == k+1 {
            res = res || backtrack(stones, j, k+1, mem)
        }
    }

    mem[key{i: i, k: k}] = res
    return res
}
func validPartition(nums []int) bool {
    return isValid(nums, len(nums), make(map[int]bool))
}

func isValid(nums []int, length int, mem map[int]bool) bool {
    if res, ok := mem[length]; ok {
        return res
    }

    size := len(nums)
    if length == 0 {
        mem[length] = true
        return true
    }

    if length == 1 {
        mem[length] = false
        return false
    }

    if length == 2 {
        mem[length] = nums[size-1] == nums[size-2]
        return mem[length]
    }

    a, b, c := nums[size-length], nums[size-(length-1)], nums[size-(length-2)]
    res := (isSame3(a,b,c) || isConsecutiveIncreasing(a,b,c)) && isValid(nums, length-3, mem)
    res = res || (a == b && isValid(nums, length-2, mem))
    mem[length] = res
    return mem[length]
}

func isSame3(a, b, c int) bool {
    return a == b && b == c
}

func isConsecutiveIncreasing(a, b, c int) bool {
    return a+1 == b && b+1 == c
}
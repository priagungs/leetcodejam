func maxSlidingWindow(nums []int, k int) []int {
    deque, res := make([]int, 0), make([]int, 0)

    for i := 0; i < k; i++ {
        deque = addToDeque(deque, nums[i], len(nums))
    }
    res = append(res, deque[0])
    if nums[0] == deque[0] {
        deque = deque[1:]
    }

    for i := 1; i < len(nums)-k+1; i++ {
        deque = addToDeque(deque, nums[i+k-1], len(nums))
        res = append(res, deque[0])
        if nums[i] == deque[0] {
            deque = deque[1:]
        }
    }
    return res
}

func addToDeque(deque []int, val, maxSize int) []int {
    for len(deque) > 0 && deque[len(deque)-1] < val {
        deque = deque[:len(deque)-1]
    }

    if len(deque) >= maxSize {
        deque = deque[1:]
    }
    return append(deque, val)
}
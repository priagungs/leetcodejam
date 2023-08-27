type key struct {
    lenS1, lenS2, lenS3 int
}

func isInterleave(s1 string, s2 string, s3 string) bool {
    return solve(s1, s2, s3, make(map[key]bool, len(s1)+len(s2)+len(s3)))
}

func solve(s1, s2, s3 string, mem map[key]bool) bool {
    memKey := key{lenS1: len(s1), lenS2: len(s2), lenS3: len(s3)}
    if val, found := mem[memKey]; found {
        return val
    }

    if s1 == s2 && s2 == s3 && s3 == "" {
        mem[memKey] = true
        return true
    }

    res := false
    for i := 0; i < len(s1) && i < len(s3) && s1[i] == s3[i]; i++ {
        res = res || solve(s1[i+1:], s2, s3[i+1:], mem)
    }

    for i := 0; i < len(s2) && i < len(s3) && s2[i] == s3[i]; i++ {
        res = res || solve(s1, s2[i+1:], s3[i+1:], mem)
    }

    mem[memKey] = res
    return res
}
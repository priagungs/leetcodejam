/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    return generate(1, n)
}

func generate(start, end int) []*TreeNode {
    if start > end {return []*TreeNode{nil}}
    if start == end {
        return []*TreeNode{
            &TreeNode{Val: start},
        }
    }

    res := make([]*TreeNode, 0)
    for i := start; i <= end; i++ {
        subTreesLeft := generate(start, i-1)
        subTreesRight := generate(i+1, end)

        for _, left := range subTreesLeft {
            for _, right := range subTreesRight {
                res = append(res, &TreeNode{
                    Val: i,
                    Left: left,
                    Right: right,
                })
            }
        }
    }
    return res
}
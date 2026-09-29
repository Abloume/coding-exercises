// Package histogram 柱状图中最大的矩形（LeetCode 84）。
//
// 给定 n 个非负整数表示柱状图中每根柱子的高度，求该柱状图中
// 能够勾勒出的矩形的最大面积。
package histogram

// LargestRectangleArea 单调递增栈 + 哨兵解法，时间复杂度 O(n)。
//
// 核心洞察：以第 i 根柱子高度为矩形高的最大面积 =
//   height[i] × (右边第一个更矮的下标 - 左边第一个更矮的下标 - 1)
//
// 单调**递增**栈（栈底到栈顶高度递增）：遍历到更矮的柱子时，
// 栈顶柱子的"右边界"确定（就是当前这根更矮的），弹出结算：
//   高 = 被弹出柱子的高度；宽 = 当前下标 - 新栈顶 - 1
// 新栈顶恰好是"左边第一个比它矮的柱子"。
//
// 哨兵：数组头尾各补一个高度 0——
//   头 0：栈永不空，任何柱子都有左边界，省掉"栈空"分支；
//   尾 0：保证遍历结束时所有柱子都被弹出结算。
//
// 与接雨水（42）的镜像关系：42 是单调**递减**栈找"两边高中间低"的
// 凹槽（水面取 min(左,右)）；本题是单调**递增**栈找"两边低中间高"的
// 柱子（矩形高取栈顶自身）。一减一增，结算逻辑正好反过来。
func LargestRectangleArea(heights []int) int {
	// 头尾补 0 哨兵：头 0 保证栈非空（左边界永远存在），尾 0 保证全部弹出。
	// 【前端转 Go 学习点】heights... 是"切片展开"语法（类比 JS 的 ...heights）：
	// 把切片元素逐个拆开作为 append 的多个参数，等价于 append([]int{0}, heights[0], heights[1], ...)。
	// 只能用于最后一个参数；只能展开切片（数组要先 arr[:]）。本行第一个参数是新切片容量仅 1，
	// 展开追加必然扩容 → 结果是全新切片，原 heights 不被修改。
	h := append([]int{0}, heights...)
	h = append(h, 0)

	stack := make([]int, 0, len(h)) // 单调递增栈，存下标
	maxArea := 0

	for i, height := range h {
		// 当前柱子比栈顶矮 → 栈顶的右边界确定，弹出结算
		for len(stack) > 0 && height < h[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 高 = 弹出柱子高度；宽 = i - 新栈顶 - 1（新栈顶即左边第一个更矮）
			area := h[top] * (i - stack[len(stack)-1] - 1)
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}

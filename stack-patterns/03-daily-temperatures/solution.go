// Package dailytemperatures 每日温度（LeetCode 739）。
//
// 给定每日温度数组，返回 answer[i]：第 i 天之后需要等多少天才会出现更高温度，
// 若之后不会升高则为 0。
package dailytemperatures

// DailyTemperatures 用单调递减栈求解，时间复杂度 O(n)。
//
// 思路：
//   - 维护一个"温度递减"的栈，栈里存的是**下标**（不是温度，因为要算距离）；
//   - 遍历到第 i 天时，只要栈顶那天的温度 < 当前温度，说明栈顶那天的
//     "下一个更高温度"就是今天 → 弹出并结算 ans[top] = i - top；
//   - 相等不算更高（严格 > 才弹出），所以栈是严格递减的；
//   - 遍历结束后留在栈里的下标，都找不到更高温度，ans 保持 0。
//
// 关键直觉：单调栈里"弹出栈顶的那一刻，它的答案确定"——因为弹出的原因就是
// 遇到了第一个比它大的元素，而这个元素就是它右侧最近的更大值。
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	// ans: 定长结果数组，用 make([]int, n) → len=n, cap=n，n 个元素初始化为零值 0，
	// 可直接按索引赋值 ans[top] = i - top。
	// 【前端转 Go 学习点】Go 切片的 len（可见长度）与 cap（底层容量）分离：
	// make([]int, n) ≈ JS 的 new Array(n).fill(0)（定长、有占位、可按下标写）。
	ans := make([]int, n)

	// stack: 动态栈，用 make([]int, 0, n) → len=0, cap=n，
	// 只预分配底层数组、不占长度，必须用 append 追加（len=0 时按下标写会越界 panic）；
	// 预分配容量可避免反复扩容时整体拷贝。
	// 【前端转 Go 学习点】Go 把"分配空间"和"放入元素"分开，JS 无对应物——
	// 这正是 len/cap 分离的核心，也是切片比数组灵活的原因。
	stack := make([]int, 0, n) // 单调递减栈，存下标

	for i, t := range temperatures {
		// 当前温度高于栈顶那天的温度 → 结算栈顶
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}

	return ans
}

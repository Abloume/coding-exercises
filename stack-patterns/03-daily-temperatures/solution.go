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
	ans := make([]int, n)
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

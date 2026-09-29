// Package rainwater 接雨水（LeetCode 42）。
//
// 给定 n 个非负整数表示宽度为 1 的柱子高度，计算下雨后能接多少雨水。
package rainwater

// Trap 用单调递减栈求解，时间复杂度 O(n)。
//
// 思路（按"层"结算水量）：
//   - 栈内存**下标**，从栈底到栈顶高度严格递减；
//   - 遍历到下标 i，只要当前高度 h 大于栈顶高度，栈顶就"凹陷"了：
//     弹出栈顶 top 作为凹槽底部，此时凹槽的
//       左边界 = 弹出后的新栈顶 left（若栈空则没有左边界，无法积水，跳出）
//       右边界 = 当前 i
//     水量 = (min(height[left], h) - height[top]) * (i - left - 1)
//   - 每个凹槽在弹出时按一层结算，凹槽内部多层由多次弹出累积。
//
// 与 739 / 496 的对比：同是"弹出栈顶的那一刻答案确定"，
// 但结算内容从"距离"（739）、"更大值"（496）变成了"水量"（本体的面积）。
func Trap(height []int) int {
	stack := make([]int, 0, len(height)) // 单调递减栈，存下标
	total := 0

	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 弹出后栈空：左边没有柱子能兜住水
			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1] // 左边界（弹出后的新栈顶）
			width := i - left - 1
			// 水量 = 两边较低者 - 凹槽底部高度，再乘宽度
			water := (min(height[left], h) - height[top]) * width
			total += water
		}
		stack = append(stack, i)
	}

	return total
}

// Package nextgreater 下一个更大元素 I（LeetCode 496）。
//
// 给定两个无重复元素的数组 nums1、nums2，nums1 是 nums2 的子集。
// 返回 nums1 中每个元素在 nums2 中右侧第一个比它大的值；不存在则为 -1。
package nextgreater

// NextGreaterElement 两段式解法：先预处理 nums2，再查询 nums1。
//
// 预处理：单调递减栈扫一遍 nums2，建"值 -> 下一个更大值"映射。
//   - 当前元素 x 大于栈顶时，栈顶的下一个更大值就是 x → 弹出并记录；
//   - 留在栈里的元素右侧无更大值 → 记为 -1。
//
// 查询：遍历 nums1，从映射取值（map 默认零值 0，若漏记会出错，
// 所以栈里剩余元素必须显式记为 -1）。
//
// 与每日温度（739）的对比：739 存下标（要算距离），本题存值（nums2 无重复，
// 值可直接当 key，且查询对象是 nums1 的"值"而非"位置"）。
//
// 时间复杂度 O(len(nums1)+len(nums2))，空间复杂度 O(len(nums2))。
func NextGreaterElement(nums1, nums2 []int) []int {
	// 【前端转 Go 学习点】map[int]int 读不存在的 key 返回零值 0 而不报错——
	// 类似 JS 的 obj[key] 得到 undefined，但 Go 不会区分"不存在"和"值为 0"，
	// 所以依赖"是否存在"时必须用双返回值 ok 或显式初始化。
	next := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2))

	for _, x := range nums2 {
		for len(stack) > 0 && x > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[top] = x // 弹出栈顶的那一刻，答案确定
		}
		stack = append(stack, x)
	}
	for _, v := range stack {
		next[v] = -1 // 剩余元素无更大值
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		ans[i] = next[v]
	}
	return ans
}

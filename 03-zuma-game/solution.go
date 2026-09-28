// Package zuma 简易祖玛消消乐。
//
// 规则：给定球序列字符串，当出现 >=3 个连续相同球时整段消除；
// 消除后两侧球合拢，若再次形成 >=3 个连续相同球则继续消除（连锁反应），
// 直到无法消除，返回最终剩余的球序列。
package zuma

// Eliminate 返回经过全部消除（含连锁）后剩余的球序列。
//
// 思路：把连续相同球压缩成 run（字符, 个数），配合栈处理。
//
// 为什么不能逐字符入栈：例如 "aabbbbaa" 的 4 个 'b'，逐字符处理时
// 第 3 个 'b' 触发消除后，第 4 个 'b' 会被当成新块重新入栈，
// 4 个 'b' 被拆成 3+1，结果错误（正确应整块消除后两侧 'a' 合拢成 4 个再消除）。
//
// 先压缩成 run 再入栈，整个连续块作为一个整体参与消除，连锁反应由
// "新 run 与栈顶合并后循环检查 >=3" 自然触发。
//
// 时间复杂度 O(n)，空间复杂度 O(n)。
func Eliminate(s string) string {
	type item struct {
		ch    byte
		count int
	}

	stack := make([]item, 0, len(s))

	for i := 0; i < len(s); {
		// 压缩当前连续相同字符块为一个 run
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		ch, cnt := s[i], j-i

		// 与栈顶同字符则合并计数，否则压入新块
		if n := len(stack); n > 0 && stack[n-1].ch == ch {
			stack[n-1].count += cnt
		} else {
			stack = append(stack, item{ch, cnt})
		}

		// 连锁消除：栈顶计数 >=3 就弹出，弹出后继续检查新栈顶
		for len(stack) > 0 && stack[len(stack)-1].count >= 3 {
			stack = stack[:len(stack)-1]
		}

		i = j
	}

	// 拼回剩余球序列
	res := make([]byte, 0, len(s))
	for _, it := range stack {
		for k := 0; k < it.count; k++ {
			res = append(res, it.ch)
		}
	}
	return string(res)
}

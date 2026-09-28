// Package genparentheses 括号生成。
//
// 给定 n 对括号，生成所有可能的且有效的括号组合。
package genparentheses

// GenerateParenthesis 生成 n 对括号的所有合法组合。
//
// 思路：回溯（DFS）+ 剪枝。
// 维护已用的左括号数 open、右括号数 close，两条剪枝规则：
//  1. open < n 时才能放 '('；
//  2. close < open 时才能放 ')'（右括号数永远不能超过左括号数，否则前缀非法）。
//
// 终止条件：当前路径长度 == 2n，记录结果。
//
// 结果数 = 第 n 个卡特兰数 C(2n,n)/(n+1)，如 n=3 时 5 种。
//
// 时间复杂度 O(4^n/√n)，空间复杂度 O(n)（不含结果）。
func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	cur := make([]byte, 0, 2*n) // 当前路径，预分配容量

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(cur) == 2*n { // 左右括号都用完
			// string(cur) 会拷贝底层数组，结果互不共享，安全
			res = append(res, string(cur))
			return
		}
		if open < n {
			cur = append(cur, '(')
			backtrack(open+1, close)
			cur = cur[:len(cur)-1] // 撤销选择
		}
		if close < open {
			cur = append(cur, ')')
			backtrack(open, close+1)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(0, 0)
	return res
}

// IsValid 判断括号字符串是否有效（供测试校验生成结果，也可独立使用）。
// 规则：'(' 入栈、')' 出栈，任何时刻右括号不得多于左括号，最终栈为空。
func IsValid(s string) bool {
	balance := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			balance++
		case ')':
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}

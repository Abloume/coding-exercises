# 01 括号生成

## 题目

给定 `n` 对括号，生成所有**可能的且有效的**括号组合。

**示例**：`n = 3` → `["((()))", "(()())", "(())()", "()(())", "()()()"]`

## 思路：回溯（DFS）+ 剪枝

把问题看成**决策树**：每一步选择放 `(` 还是 `)`，用 `open` / `close` 记录已用的左右括号数。

两条剪枝规则是核心：

1. `open < n` 时才能放 `(`——左括号总共只有 n 个；
2. `close < open` 时才能放 `)`——**任意前缀中右括号数不能超过左括号数**，否则一定非法。

```go
func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	cur := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(cur) == 2*n {          // 左右括号都用完 → 记录结果
			res = append(res, string(cur))
			return
		}
		if open < n {                  // 剪枝 1：还能放左括号
			cur = append(cur, '(')
			backtrack(open+1, close)
			cur = cur[:len(cur)-1]    // 撤销选择（回溯）
		}
		if close < open {              // 剪枝 2：右括号少于左括号时才能放
			cur = append(cur, ')')
			backtrack(open, close+1)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(0, 0)
	return res
}
```

### 常见的三个坑

1. **只剪枝 1 不剪枝 2**：会生成 `"())("` 这类非法串。剪枝 2（`close < open`）保证**每一步前缀都合法**，这是"生成合法括号"的关键。
2. **终止条件写错**：应该按 `len(cur) == 2*n` 判断，而不是 `close == n`（右括号数等于 n 但左括号未用完时路径未结束）。
3. **Go 切片回溯陷阱**：`cur` 的 append / 裁剪都在**同一个底层数组**上操作。这里 `string(cur)` 会**拷贝**一份底层数据，所以存入 `res` 的结果是独立的；如果直接存 `cur`（如 `res = append(res, cur)`），后续回溯会**覆盖**前面记录的结果，得到一堆相同字符串。这正是"回溯撤销要用拷贝"的典型场景。

## 复杂度

- 结果数 = 第 n 个卡特兰数 `C(2n, n)/(n+1)`（n=3 是 5，n=4 是 14，n=5 是 42）
- 时间：O(4ⁿ/√n)；空间：O(n)（递归深度 + 路径长度，不含结果集）

## 变种题（大厂高频，务必顺带掌握）

| 变种 | 题号 | 思路 | 备注 |
|---|---|---|---|
| 有效括号（判断） | LeetCode 20 | 栈 / 计数器，见本练习 `IsValid` | 最基础，腾讯一面常考 |
| 最长有效括号 | LeetCode 32 | 栈存下标 或 DP | hard 档，字节/虾皮可能出现 |
| 最小添加使括号有效 | LeetCode 921 | 贪心计数，`balance` 为负时补左括号 | 简单，可作热身 |
| 移除无效的括号 | LeetCode 301 | BFS 逐层删 或 DFS+剪枝 | hard，考察剪枝功底 |
| 括号的分数 | LeetCode 856 | 栈 / 递归，内层分数×2 | 与本题思路互补 |
| 有效括号字符串（含 `*`） | LeetCode 678 | 双计数器（lo/hi 区间） | 变体题，考察扩展思维 |
| 卡特兰数家族 | 出栈序列 / 不同 BST / 凸多边形三角划分 | 公式或 DP | 数量类变种，常被追问 |

## 关联考点

- 回溯模板：选择 → 递归 → 撤销（选择列表、路径、终止条件三要素）
- 剪枝的本质：提前排除不可能的分支，比"生成后校验"高效得多
- 卡特兰数：`f(n) = Σ f(i)·f(n-1-i)`，和二叉树计数、出栈序列同源

## 运行测试

```bash
# 只测当前题目
cd backtracking-patterns/01-generate-parentheses && go test -v ./...

# 测全部题目（项目根目录）
for d in */*/; do (cd "$d" && go test ./...); done
```

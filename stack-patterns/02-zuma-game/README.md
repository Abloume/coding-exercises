# 02 祖玛消消乐

## 题目

给定一个球序列字符串（如 `"RRBBBR"`），当出现 **>=3 个连续相同球** 时，整段消除；消除后两侧球合拢，若再次形成 >=3 个连续相同球，继续消除（**连锁反应**），直到无法消除。返回最终剩余的球序列。

**示例**：`"aabbbbaa"` → 消 `bbbb` → `"aaaa"` → 消 `aaaa` → `""`（空）

## 思路：run 压缩 + 栈

把连续相同球**先压缩成一个 run（字符, 个数）**，再配合栈处理：

1. 遍历时跳过连续相同字符，得到 run；
2. run 与栈顶同字符则合并计数，否则压入新块；
3. **循环检查**栈顶计数 >=3 就弹出——弹出后继续检查新栈顶，连锁反应由此触发。

```go
func Eliminate(s string) string {
	type item struct {
		ch    byte
		count int
	}
	stack := make([]item, 0, len(s))

	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] { // 压缩 run
			j++
		}
		ch, cnt := s[i], j-i

		if n := len(stack); n > 0 && stack[n-1].ch == ch {
			stack[n-1].count += cnt        // 与栈顶合并
		} else {
			stack = append(stack, item{ch, cnt})
		}

		for len(stack) > 0 && stack[len(stack)-1].count >= 3 {
			stack = stack[:len(stack)-1]   // 连锁消除
		}
		i = j
	}

	res := make([]byte, 0, len(s))
	for _, it := range stack {
		for k := 0; k < it.count; k++ {
			res = append(res, it.ch)
		}
	}
	return string(res)
}
```

## 本题最大的坑：为什么不能逐字符入栈

以 `"aabbbbaa"` 为例（4 个连续 `b`）：

- **逐字符入栈**：第 3 个 `b` 让栈顶 `(b,3)` 触发消除弹出，第 4 个 `b` 被当成**新块**重新入栈 → 4 个 `b` 被拆成 3+1，结果错误（`"aabbaa"`）；
- **run 压缩入栈**：4 个 `b` 作为 `(b,4)` 一个整体入栈，一次弹出；两侧 `a` 合拢后与后续 run 合并成 `(a,4)`，再次弹出 → 正确得到空串。

连锁反应的正确形态：**新 run 到达时与栈顶合并 → 检查 >=3 → 弹出 → 新栈顶可能又 >=3 → 继续弹**。run 压缩保证"同一个连续块"不会被拆分。

## 复杂度

- 时间：O(n)（每个字符最多入栈/出栈一次）
- 空间：O(n)（栈）

## 变体题

| 变体 | 说明 | 备注 |
|---|---|---|
| LeetCode 488 祖玛游戏（完整版） | 给定桌面球与手中球，求**最少**用几颗手中球能清空桌面 | hard，DFS + 记忆化/剪枝，一面较少见但字节问过 |
| 消除后求剩余（本练习） | 只做消除，返回最终串 | 字节/滴滴高频，栈题 |
| 连续 k 次消除 | 阈值不是 3 而是 k | 参数化 run 计数即可，思路相同 |

## 关联考点

- run-length 压缩：`O(n)` 扫描连续相同段的通用技巧
- 栈顶合并 + 循环弹出的模式，也用于"合并区间""每日温度"等单调栈变体
- 注意 Go 中 `item{ch, cnt}` 的匿名字段结构体切片用法

## 运行测试

```bash
# 只测当前题目
cd stack-patterns/02-zuma-game && go test -v ./...

# 测全部题目（项目根目录）
for d in */*/; do (cd "$d" && go test ./...); done
```

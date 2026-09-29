# 04 下一个更大元素 I

> 模式：栈 **C1 单调栈 · 下一个更大元素**（见 [STACK-PATTERNS.md](../../STACK-PATTERNS.md)）

## 题目

给定两个**没有重复元素**的数组 `nums1` 和 `nums2`（`nums1` 是 `nums2` 的子集）。返回 `nums1` 中每个元素在 `nums2` 中**右侧第一个比它大的值**，不存在则为 `-1`。

**示例**：`nums1 = [4,1,2]`，`nums2 = [1,3,4,2]` → `[-1, 3, -1]`

## 思路：预处理（单调栈）+ 查询（map）

两段式：

**第一段**：单调递减栈扫一遍 `nums2`，建立 `值 → 下一个更大值` 的映射。
**第二段**：遍历 `nums1`，从 map 里查答案。

```go
func NextGreaterElement(nums1, nums2 []int) []int {
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
```

### 与每日温度（739）的对比（同族关键差异）

| | 每日温度 739 | 下一个更大元素 I 496 |
|---|---|---|
| 单调栈里存什么 | **下标**（要算天数差） | **值**（nums2 无重复，值可当 key） |
| 结算内容 | `ans[top] = i - top`（距离） | `next[top] = x`（更大值本身） |
| 查询对象 | 原数组按位置 | nums1 按值查 map |
| 本质 | 记录位置 → 算距离 | 记录值 → 建映射 |

**为什么本题可以存值**：题目保证 nums2 无重复 → 值唯一 → 能直接做 map key。如果允许重复，就得回到存下标。

### 两个坑

1. **map 读不存在的 key 返回 0 而不是报错**：Go 的 `next[v]` 在 key 不存在时返回零值——所以栈里剩余元素必须**显式** `next[v] = -1`，否则它们会被查成 0；
2. **不要对每个 nums1 元素去 nums2 里重新找**：那是 O(n×m)，预处理一次 O(n) 扫完才是单调栈的意义。

## 复杂度

- 时间：O(len(nums1) + len(nums2))（两段各扫一遍）
- 空间：O(len(nums2))（map + 栈）

## 变体题（同族）

| 变体 | 题号 | 变化 |
|---|---|---|
| 下一个更大元素 II | 503 | 循环数组：遍历 `2 * n` 次，下标用 `i % n` |
| 每日温度 | 739 | 已练：存下标算距离（本题的"前身"） |
| 接雨水 | 42 | 同一单调递减栈思路，结算水量 |

## 运行测试

```bash
cd stack-patterns/04-next-greater-element-i && go test -v ./...
```

# coding-exercises

日常编程练习题集，使用 Go 实现。

## 练习列表

按**模式**组织文件夹（编号 = 文件夹内递增；主模式 = 该题最优解所用的模式）：

| 模式 | 编号 | 题目 | 目录 |
|---|---|---|---|
| 栈 | 01 | 路径简化 | [stack-patterns/01-path-simplify](./stack-patterns/01-path-simplify) |
| 栈 | 02 | 祖玛消消乐 | [stack-patterns/02-zuma-game](./stack-patterns/02-zuma-game) |
| 栈 | 03 | 每日温度 | [stack-patterns/03-daily-temperatures](./stack-patterns/03-daily-temperatures) |
| 栈 | 04 | 下一个更大元素 I | [stack-patterns/04-next-greater-element-i](./stack-patterns/04-next-greater-element-i) |
| 栈 | 05 | 接雨水 | [stack-patterns/05-trapping-rain-water](./stack-patterns/05-trapping-rain-water) |
| 栈 | 06 | 柱状图中最大的矩形 | [stack-patterns/06-largest-rectangle-in-histogram](./stack-patterns/06-largest-rectangle-in-histogram) |
| 回溯 | 01 | 括号生成 | [backtracking-patterns/01-generate-parentheses](./backtracking-patterns/01-generate-parentheses) |

> 模式索引见 [STACK-PATTERNS.md](./STACK-PATTERNS.md)。

## 说明

每个练习为独立 Go module，含题目说明（README.md）、解法（solution.go）与测试（solution_test.go）。
根目录的 `go.work` 将各练习模块纳入同一个 Go workspace，可跨目录共享依赖解析。

## 运行测试

**只测某个练习**：`cd` 到该练习目录再 `go test`（以 01 路径简化为准）：

```bash
cd stack-patterns/01-path-simplify && go test -v ./...
```

**测全部练习**（在项目根目录）：

```bash
for d in */*/; do (cd "$d" && go test ./...); done
```

> 注：因为各练习是独立 module，根目录直接 `go test ./...` 不可用（根目录不属于任何 module），
> 全部题目用上面的循环方式。

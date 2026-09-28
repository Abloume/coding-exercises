// Package pathsimplify 将路径化简为最简形式。
//
// 规则：仅 "." 表示当前目录、".." 表示上一级目录，
// 其余分段（包括 "...." 这类特殊名字）一律视为普通目录名。
package pathsimplify

import "strings"

// SimplifyPath 将 path 化简为最终路径（相对路径形式）。
//
// 用栈模拟目录进出：遇到 ".." 且栈非空时出栈，
// 栈为空（已在根）时 ".." 无效果；"." 与空段直接忽略。
//
// 时间复杂度 O(n)，空间复杂度 O(n)，n 为路径长度。
func SimplifyPath(path string) string {
	stack := make([]string, 0, 8)
	for _, seg := range strings.Split(path, "/") {
		switch seg {
		case "", ".":
			// 空段（连续 "/"）与当前目录：忽略
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			// 已在根时 ".." 停在原地，不做任何事
		default:
			stack = append(stack, seg)
		}
	}
	return strings.Join(stack, "/")
}

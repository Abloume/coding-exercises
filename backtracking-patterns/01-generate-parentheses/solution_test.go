package genparentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n=1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n=2",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "n=3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesis(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// TestCatalanCount 校验生成数量符合卡特兰数：n=4 应为 14，n=5 应为 42。
func TestCatalanCount(t *testing.T) {
	catalan := []int{1, 1, 2, 5, 14, 42} // 下标即 n
	for n, want := range catalan {
		if n == 0 {
			continue
		}
		if got := len(GenerateParenthesis(n)); got != want {
			t.Errorf("GenerateParenthesis(%d) count = %d, want %d", n, got, want)
		}
	}
}

// TestAllValid 校验每个生成结果都满足括号有效性，且无重复。
func TestAllValid(t *testing.T) {
	for n := 1; n <= 6; n++ {
		res := GenerateParenthesis(n)
		seen := make(map[string]bool, len(res))
		for _, s := range res {
			if !IsValid(s) {
				t.Fatalf("GenerateParenthesis(%d) produced invalid: %q", n, s)
			}
			if len(s) != 2*n {
				t.Fatalf("GenerateParenthesis(%d) wrong length: %q", n, s)
			}
			if seen[s] {
				t.Fatalf("GenerateParenthesis(%d) duplicated: %q", n, s)
			}
			seen[s] = true
		}
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()()", true},
		{"(())", true},
		{"())", false},
		{"(()", false},
		{")(()", false},
	}
	for _, tt := range tests {
		if got := IsValid(tt.s); got != tt.want {
			t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

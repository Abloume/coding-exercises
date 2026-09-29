package histogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "classic leetcode case",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "two bars",
			heights: []int{2, 4},
			want:    4,
		},
		{
			name:    "equal heights",
			heights: []int{1, 1},
			want:    2,
		},
		{
			name:    "single bar",
			heights: []int{5},
			want:    5,
		},
		{
			name:    "valley shape",
			heights: []int{3, 1, 3},
			want:    3,
		},
		{
			name:    "increasing",
			heights: []int{1, 2, 3, 4, 5},
			want:    9,
		},
		{
			name:    "decreasing",
			heights: []int{5, 4, 3, 2, 1},
			want:    9,
		},
		{
			name:    "peak in middle",
			heights: []int{2, 1, 2},
			want:    3,
		},
		{
			name:    "all zero",
			heights: []int{0, 0, 0},
			want:    0,
		},
		{
			name:    "empty",
			heights: []int{},
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("LargestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}

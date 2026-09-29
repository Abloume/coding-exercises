package rainwater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "classic leetcode case",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "classic case 2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "small basin",
			height: []int{1, 0, 1},
			want:   1,
		},
		{
			name:   "two walls",
			height: []int{3, 0, 3},
			want:   3,
		},
		{
			name:   "two walls height 2",
			height: []int{2, 0, 2},
			want:   2,
		},
		{
			name:   "decreasing no water",
			height: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "increasing no water",
			height: []int{1, 2, 3, 4},
			want:   0,
		},
		{
			name:   "single element",
			height: []int{0},
			want:   0,
		},
		{
			name:   "empty",
			height: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

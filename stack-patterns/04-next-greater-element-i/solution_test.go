package nextgreater

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "classic leetcode case",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			name:  "monotonic increasing nums2",
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
		{
			name:  "all resolved by last element",
			nums1: []int{1, 3, 5, 2, 4},
			nums2: []int{6, 5, 4, 3, 2, 1, 7},
			want:  []int{7, 7, 7, 7, 7},
		},
		{
			name:  "single element no greater",
			nums1: []int{2},
			nums2: []int{1, 2},
			want:  []int{-1},
		},
		{
			name:  "last element itself queried",
			nums1: []int{1, 2},
			nums2: []int{1, 2},
			want:  []int{2, -1},
		},
		{
			name:  "empty nums1",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  []int{},
		},
		{
			name:  "all have next greater",
			nums1: []int{1, 2},
			nums2: []int{1, 3, 2, 4},
			want:  []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterElement(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElement(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}

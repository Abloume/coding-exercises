package dailytemperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name string
		t    []int
		want []int
	}{
		{
			name: "classic leetcode case",
			t:    []int{73, 74, 75, 71, 69, 72, 76, 73},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name: "strictly increasing",
			t:    []int{30, 40, 50, 60},
			want: []int{1, 1, 1, 0},
		},
		{
			name: "strictly decreasing",
			t:    []int{90, 80, 70, 60},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "equal temps not higher",
			t:    []int{30, 30, 30},
			want: []int{0, 0, 0},
		},
		{
			name: "single element",
			t:    []int{73},
			want: []int{0},
		},
		{
			name: "empty",
			t:    []int{},
			want: []int{},
		},
		{
			name: "peak in the middle",
			t:    []int{55, 50, 60, 40, 70},
			want: []int{2, 1, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DailyTemperatures(tt.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DailyTemperatures(%v) = %v, want %v", tt.t, got, tt.want)
			}
		})
	}
}

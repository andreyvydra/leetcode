package tasks

import "testing"

func TestIsRectangleOverlap(t *testing.T) {
	tests := []struct {
		name string
		rec1 []int
		rec2 []int
		want bool
	}{
		{
			name: "broken case",
			rec1: []int{7, 8, 13, 15},
			rec2: []int{10, 8, 12, 20},
			want: true,
		},
		{
			name: "right top overlap",
			rec1: []int{0, 0, 2, 2},
			rec2: []int{1, 1, 3, 3},
			want: true,
		},
		{
			name: "common side",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{1, 0, 2, 1},
			want: false,
		},
		{
			name: "no overlap",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{2, 2, 3, 3},
			want: false,
		},
		{
			name: "full overlap by 2",
			rec1: []int{5, 5, 6, 6},
			rec2: []int{1, 1, 10, 10},
			want: true,
		},
		{
			name: "full overlap by 1",
			rec1: []int{1, 1, 10, 10},
			rec2: []int{5, 5, 6, 6},
			want: true,
		},
		{
			name: "1/4 from rec1",
			rec1: []int{4, 4, 6, 6},
			rec2: []int{5, 5, 6, 6},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRectangleOverlap(tt.rec1, tt.rec2)
			if got != tt.want {
				t.Errorf("isRectangleOverlap(%v, %v) = %v; want %v", tt.rec1, tt.rec2, got, tt.want)
			}
		})
	}
}

package tasks

import "testing"

func TestLargestOverlap(t *testing.T) {
	tests := []struct {
		name string
		img1 [][]int
		img2 [][]int
		want int
	}{
		{
			name: "simple 3x3",
			img1: [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			img2: [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			want: 3,
		},
		{
			name: "1x1",
			img1: [][]int{{1}},
			img2: [][]int{{1}},
			want: 1,
		},
		{
			name: "1x1 (zero)",
			img1: [][]int{{0}},
			img2: [][]int{{0}},
			want: 0,
		},
		{
			name: "round 4x4",
			img1: [][]int{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}},
			img2: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 1}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestOverlap(tt.img1, tt.img2)
			if got != tt.want {
				t.Errorf("largestOverlap(%v, %v) = %d; want %d", tt.img1, tt.img2, got, tt.want)
			}
		})
	}
}

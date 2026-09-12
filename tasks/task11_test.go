package tasks

import "testing"

func TestMaxArea(t *testing.T) {
	result := maxArea([]int{4, 4, 2, 4})
	if result != 12 {
		t.Errorf("maxArea = %d; expected 12", result)
	}
}

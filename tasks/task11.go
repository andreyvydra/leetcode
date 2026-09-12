package tasks

func maxArea(height []int) int {
	// l * min(h1, h2)

	l, r := 0, len(height)-1
	m := 0
	for l < r {
		m = max(m, (r-l)*min(height[r], height[l]))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return m
}

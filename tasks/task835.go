package tasks

func largestOverlap(img1 [][]int, img2 [][]int) int {
	setImg2 := make(map[[2]int]struct{})
	setImg1 := make(map[[2]int]struct{})

	for i, row := range img2 {
		for j, cell := range row {
			if cell == 1 {
				setImg2[[2]int{i, j}] = struct{}{}
			}
			if img1[i][j] == 1 {
				setImg1[[2]int{i, j}] = struct{}{}
			}
		}
	}

	m := 0
	diffs := make(map[[2]int]int)
	for cord1 := range setImg1 {
		for cord2 := range setImg2 {
			diff1 := cord2[0] - cord1[0]
			diff2 := cord2[1] - cord1[1]
			diffs[[2]int{diff1, diff2}] += 1
			m = max(diffs[[2]int{diff1, diff2}], m)
		}
	}

	return m
}

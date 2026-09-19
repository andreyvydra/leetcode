package tasks

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	closestX := 0
	if x1 <= xCenter && x2 >= xCenter {
		closestX = xCenter
	} else if (xCenter-x1)*(xCenter-x1) > (xCenter-x2)*(xCenter-x2) {
		closestX = x2
	} else {
		closestX = x1
	}

	closestY := 0
	if y1 <= yCenter && y2 >= yCenter {
		closestY = yCenter
	} else if (yCenter-y1)*(yCenter-y1) > (yCenter-y2)*(yCenter-y2) {
		closestY = y2
	} else {
		closestY = y1
	}

	dx := xCenter - closestX
	dy := yCenter - closestY

	return dx*dx+dy*dy <= radius*radius

}

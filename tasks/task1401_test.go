package tasks

import (
	"testing"
)

func TestCheckOverlap(t *testing.T) {
	tests := []struct {
		name    string
		radius  int
		xCenter int
		yCenter int
		x1      int
		y1      int
		x2      int
		y2      int
		want    bool
	}{
		{
			name:    "share point (1,0)",
			radius:  1,
			xCenter: 0,
			yCenter: 0,
			x1:      1,
			y1:      -1,
			x2:      3,
			y2:      1,
			want:    true,
		},
		{
			name:    "no share point",
			radius:  1,
			xCenter: 1,
			yCenter: 1,
			x1:      1,
			y1:      -3,
			x2:      2,
			y2:      -1,
			want:    false,
		},
		{
			name:    "left right share",
			radius:  1,
			xCenter: 0,
			yCenter: 0,
			x1:      -1,
			y1:      0,
			x2:      0,
			y2:      1,
			want:    true,
		},
		{
			name:    "square into share",
			radius:  5,
			xCenter: 0,
			yCenter: 0,
			x1:      -2,
			y1:      -2,
			x2:      2,
			y2:      2,
			want:    true,
		},
		{
			name:    "shit case",
			radius:  5,
			xCenter: 0,
			yCenter: 0,
			x1:      4,
			y1:      4,
			x2:      5,
			y2:      5,
			want:    false,
		},
		{
			name:    "shit case",
			radius:  1,
			xCenter: 1,
			yCenter: 1,
			x1:      -3,
			y1:      -3,
			x2:      3,
			y2:      3,
			want:    true,
		},
		{
			name:    "shit case 3",
			radius:  10,
			xCenter: 10,
			yCenter: 1,
			x1:      0,
			y1:      0,
			x2:      100,
			y2:      100,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkOverlap(tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2)
			if got != tt.want {
				t.Errorf("checkOverlap(%v, %v, %v, %v, %v, %v, %v) = %v; want %v", tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2, got, tt.want)
			}
		})
	}
}

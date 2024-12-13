package main

import (
	"testing"
)

func TestSumPoints(t *testing.T) {
	tests := []struct {
		a, x1, y1, x2, y2, ord, expX, expY int
	}{
		{2, 2, 2, 2, 2, 13, 5, 7},
		{2, 2, 3, 4, 4, 7, 3, 0},
		{2, 1, 11, 1, 11, 13, 2, 0},
	}
	for _, tt := range tests {
		resX, resY := sumPoints(tt.a, tt.x1, tt.y1, tt.x2, tt.y2, tt.ord)
		if resX != tt.expX || resY != tt.expY {
			t.Errorf("sumPoints(%d;%d and %d;%d ) = (%d;%d); want( %d;%d)", tt.x1, tt.y1, tt.x2, tt.y2, resX, resY, tt.expX, tt.expY)
		}
	}
}

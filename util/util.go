package util

import "C"
import "github.com/go-vgo/robotgo"

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func ImageSize(bitmap *robotgo.CBitmap) (w, h int) {
	gbit := robotgo.ToBitmap(robotgo.ToMMBitmapRef(*bitmap))
	w = gbit.Width
	h = gbit.Height
	return w, h
}
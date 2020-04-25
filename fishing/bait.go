package fishing

import (
	"fmt"
	"image"
	"math"
)

type BaitDetector interface {
	CheckBait(position *image.Point) bool
}

func CreateBaitDetector() BaitDetector {
	return newBaitDetectorImpl()
}

func newBaitDetectorImpl() BaitDetector {
	return &BaitDetectorImpl{}
}

type BaitDetectorImpl struct {
	left,top, right, bottom int
	once bool
}

func (bd *BaitDetectorImpl) CheckBait(position *image.Point) bool {
	if nil == position {
		return false
	}
	if bd.once && (math.Abs(float64(bd.left - position.X)) > 50 || math.Abs(float64(bd.top - position.Y)) > 50) {
		fmt.Print("Skipping outlier")
		return false
	}


	if !bd.once || position.X < bd.left {
		bd.left = position.X
	}
	if !bd.once || position.X > bd.right {
		bd.right = position.X
	}

	if !bd.once || position.Y < bd.top {
		bd.top = position.Y
	}
	if !bd.once || position.Y > bd.bottom {
		bd.bottom = position.Y
	}

	bd.once = true
	return bd.detect()
}

func (bd *BaitDetectorImpl) detect() bool {

	x := bd.left
	y := bd.top
	w := bd.right - bd.left
	h := bd.bottom - bd.top

	fmt.Printf("[%v, %v - %v, %v]\n", x, y, w, h)
	return w >= 15 || h >= 15
}
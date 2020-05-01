package util

type Rect interface {
	Top() int
	Bottom() int
	Left() int
	Right() int
}

type RectImpl struct {
	top    int
	bottom int
	left   int
	right  int
}

func (r *RectImpl) Top() int {
	return r.top
}

func (r *RectImpl) Bottom() int {
	return r.bottom
}

func (r *RectImpl) Left() int {
	return r.left
}

func (r *RectImpl) Right() int {
	return r.right
}

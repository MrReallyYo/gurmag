package fishing

import "github.com/go-vgo/robotgo"

func NewRobotGoScreen() *RobotGoScreen {
	return &RobotGoScreen{
		screen: robotgo.CBitmap(robotgo.CaptureScreen()),
	}
}

type RobotGoScreen struct {
	screen robotgo.CBitmap
}

func (s *RobotGoScreen) Size() (w, h int) {
	if nil == s || nil == s.screen {
		return -1, -1
	}

	bitmap := robotgo.ToBitmap(robotgo.ToMMBitmapRef(s.screen))
	w = bitmap.Width
	h = bitmap.Height
	return w, h
}

func (s *RobotGoScreen) Color(x, y int) (r, g, b int) {
	if nil == s || nil == s.screen {
		return -1, -1, -1
	}
	color := robotgo.CHex(robotgo.GetColor(robotgo.ToMMBitmapRef(s.screen), x, y))
	return int((color >> 16) & 0xFF), int((color >> 8) & 0xFF), int((color) & 0xFF)
}

func (s *RobotGoScreen) Free() {
	if nil == s || nil == s.screen {
		return
	}
	robotgo.FreeBitmap(robotgo.ToMMBitmapRef(s.screen))
	s.screen = nil
}

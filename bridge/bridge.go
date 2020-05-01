package bridge

// #include "bridge.h"
import "C"

func GetWindowPos() (x, y, w, h int64) {
	pos := C.GetWindowPos(C.CString("World of Warcraft"), C.CString("GxWindowClass"))
	//pos := C.GetWindowPos(C.CString("spotify premium"))

	return int64(pos.x), int64(pos.y), int64(pos.w), int64(pos.h)
}

func GetMousePos() (int, int) {
	pos := C.GetMousePos()
	return int(pos.x), int(pos.y)
}

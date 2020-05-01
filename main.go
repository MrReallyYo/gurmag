package main

import "coolstory.eu/gurmag/bridge"

func main() {

	x2, y2 := bridge.GetMousePos()
	println(x2)
	println(y2)

	x, y, w, h := bridge.GetWindowPos()
	println(x)
	println(y)
	println(w)
	println(h)

	/*bla, _ :=

	screen := fishing.NewRobotGoScreen()
	w, _ = screen.Size()
	println(w)
	screen.Free()

	a,b := robotgo.GetMousePos()
	println(a)
	println(b)*/
}

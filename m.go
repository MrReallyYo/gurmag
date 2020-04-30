package main

import (
	"coolstory.eu/gurmag/fishing"
	"fmt"
	"github.com/go-vgo/robotgo"
	"image"
	"time"
)


func cast() {
	robotgo.KeyTap("3")
	time.Sleep(2 * time.Second)
}

func loot(pos *image.Point) {
	robotgo.Move(pos.X, pos.Y)
	robotgo.Click("right")
}

func find(finder fishing.BobberFinder, knownPosition *image.Point) *image.Point {
	screen := robotgo.CBitmap(robotgo.CaptureScreen())
	defer robotgo.FreeBitmap(robotgo.ToMMBitmapRef(screen))
	return finder.FindBobber(&screen,nil)
}

func main() {


	finder := fishing.CreateBobberFinder()
	for true {

		cast()

		pos := find(finder, nil)
		fmt.Print(pos)
		if nil != pos {
			bait := fishing.CreateBaitDetector()
			bait.CheckBait(pos)
			Wait: for i := 0; i < 200; i++ {
				if bait.CheckBait(find(finder, pos)) {
					loot(pos)
					break Wait
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
		time.Sleep(2 * time.Second)

	}



}

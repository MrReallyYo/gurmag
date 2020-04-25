package main

import (
	"coolstory.eu/gurmag/fishing"
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {

	finder := fishing.CreateBobberFinder()

	for true {
		robotgo.KeyTap("3")
		time.Sleep(1 * time.Second)


		pos := finder.FindBobber(nil)
		fmt.Print(pos)
		if nil != pos {

			bait := fishing.CreateBaitDetector()
			bait.CheckBait(pos)
			Wait: for i := 0; i < 200; i++ {
				if bait.CheckBait(finder.FindBobber(pos)) {
					robotgo.Move(pos.X, pos.Y)
					robotgo.Click("right")
					break Wait
				}
				time.Sleep(100 * time.Millisecond)
			}
		}

		time.Sleep(2 * time.Second)

	}



}

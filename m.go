package main

import (
	"coolstory.eu/gurmag/fishing"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"time"
)

func main() {

	finder := fishing.CreateBobberFinder()


		pos := finder.FindBobber()
		fmt.Print(pos)
		if nil != pos {

			bobberMissingTicks := 0
			for i := 0; i < 500; i++ {
				if finder.CheckBait(pos) {
					fmt.Print("missing")
					bobberMissingTicks++
				} else {
					bobberMissingTicks = 0
				}
				if bobberMissingTicks >= 10 {
					fmt.Print("bait?")
					robotgo.Move(pos.X, pos.Y)
					robotgo.Click("right")
					os.Exit(1)
				}
				time.Sleep(50 * time.Millisecond)
			}


		}


}

package main

import (
	"coolstory.eu/gurmag/fishingold"
	"github.com/go-vgo/robotgo"
	"testing"
)

func TestFind1(t *testing.T) {
	runTest(t, "res/test/test.png", true)
}

func TestFind2(t *testing.T) {
	runTest(t, "res/test/bobberraw.png", true)
}

func TestFind3(t *testing.T) {
	runTest(t, "res/test/full.png", true)
}

func TestFind4(t *testing.T) {
	runTest(t, "res/test/bobberv4.png", true)
}

func TestFind5(t *testing.T) {
	runTest(t, "res/test/full2.png", true)
}

func TestFail1(t *testing.T) {
	runTest(t, "res/test/fail.png", false)
}

func runTest(t *testing.T, file string, shouldFind bool) {
	if test(file) != shouldFind {
		if shouldFind {
			t.Errorf("Should have found bobber, but failed.")
		} else {
			t.Errorf("Shouldn't have found bobber, but found.")
		}
	}
}

func test(file string) bool {
	finder := fishingold.CreateBobberFinder()
	screen := robotgo.CBitmap(robotgo.OpenBitmap(file))
	pos := finder.FindBobber(&screen, nil)
	robotgo.FreeBitmap(robotgo.ToMMBitmapRef(screen))
	return nil != pos
}

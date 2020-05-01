package fishingold

import "C"
import (
	"coolstory.eu/gurmag/util"
	"fmt"
	"github.com/go-vgo/robotgo"
	"image"
)

type BobberFinder interface {
	FindBobber(screen *robotgo.CBitmap, knownPosition *image.Point) *image.Point
}

func CreateBobberFinder() BobberFinder {
	return newBobberFinderImpl()
}

func newBobberFinderImpl() BobberFinderImpl {
	return BobberFinderImpl{}
}

type BobberFinderImpl struct {
}

func (bf BobberFinderImpl) FindBobber(screen *robotgo.CBitmap, knownPosition *image.Point) *image.Point {

	var found *image.Point

	search := func(x, y int) bool {
		found = bf.check(screen, x, y)
		if nil != found {
			fmt.Printf("bobber @ %v, %v\n", found.X, found.Y)
		}
		return nil != found
	}

	if nil != knownPosition {
		seachNearby(screen, knownPosition.X, knownPosition.Y, 50, 50, 50, 50, search)
	} else {
		w, h := util.ImageSize(screen)
		seachNearby(screen, 0, 0, 0, w, 0, h, search)
	}
	return found
}

const verbose = false

func (bf BobberFinderImpl) check(screen *robotgo.CBitmap, x int, y int) *image.Point {
	// find red starting pixel
	if isRed(screen, x, y) {
		redFeather := search(screen, x, y, 30, isRed)
		blueFeather := search(screen, x, y, 30, isBlue)
		//bobber := search(screen, x, y, 10, isBobber)
		//glow := search(screen, x, y, 5, isGlow)
		hook := search(screen, x, y, 15, isHook)
		//fmt.Printf("%v, %v, %v -  %v, %v\n", redFeather, blueFeather, hook, x, y)
		if redFeather && blueFeather && hook {
			return &image.Point{X: x, Y: y}
		}
	}

	return nil
}

func seachNearby(screen *robotgo.CBitmap, x, y, before, after, over, under int, call func(x, y int) bool) bool {

	w, h := util.ImageSize(screen)

	searchWBegin := util.Max(0, x-before)
	searchWEnd := util.Min(w, x+after)
	searchHBegin := util.Max(0, y-over)
	searchHEnd := util.Min(h, y+under)

	for searchY := searchHBegin; searchY < searchHEnd; searchY++ {
		for searchX := searchWBegin; searchX < searchWEnd; searchX++ {
			if call(searchX, searchY) {
				return true
			}
		}
	}
	return false
}

func color(screen *robotgo.CBitmap, x int, y int) (r int, g int, b int) {
	color := robotgo.CHex(robotgo.GetColor(robotgo.ToMMBitmapRef(*screen), x, y))
	return int((color >> 16) & 0xFF), int((color >> 8) & 0xFF), int((color) & 0xFF)
}

func search(screen *robotgo.CBitmap, x, y, pixelNeeded int, check func(screen *robotgo.CBitmap, x, y int) bool) bool {
	pixelFound := 0
	return seachNearby(screen, x, y, 10, 10, 10, 10, func(x int, y int) bool {
		if check(screen, x, y) {
			pixelFound++
		}
		return pixelFound >= pixelNeeded
	})
}

func isRed(screen *robotgo.CBitmap, x, y int) bool {
	r, g, b := color(screen, x, y)

	red := 1.9
	minRed := 50
	maxRed := 150
	gbMin := 0.50
	gbMax := 1.8
	gb := float64(g) / float64(b)

	return r >= minRed && r <= maxRed && r > int(float64(g)*red) && r > int(float64(b)*red) && gb >= gbMin && gb <= gbMax
}

func isBlue(screen *robotgo.CBitmap, a, c int) bool {
	r, g, b := color(screen, a, c)

	blue := 1.1
	minBlue := 60

	rgMin := 0.8
	rgMax := 1.2
	rg := float64(r) / float64(g)

	return b >= minBlue && b > int(float64(r)*blue) && b > int(float64(g)*blue) && g > r && rg >= rgMin && rg <= rgMax
}
func isBobber(screen *robotgo.CBitmap, a, c int) bool {
	bobber := 1.1
	minBlue := 45
	r, g, b := color(screen, a, c)
	return b >= minBlue && r > int(float64(b)*bobber) && g > int(float64(b)*bobber)
}
func isGlow(screen *robotgo.CBitmap, a, c int) bool {
	glow := 1.05
	minBlue := 55
	r, g, b := color(screen, a, c)
	return b >= minBlue && r > int(float64(b)*glow) && g > int(float64(b)*glow)
}

func isHook(screen *robotgo.CBitmap, a, c int) bool {
	minR := 70
	r, g, b := color(screen, a, c)
	return r >= minR && r <= g && r <= b
}

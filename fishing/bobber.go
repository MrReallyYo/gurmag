package fishing

import "C"
import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"image"
)

type BobberFinder interface {
	FindBobber(knownPosition *image.Point) *image.Point
}

func CreateBobberFinder() BobberFinder {
	return newBobberFinderImpl()
}

func newBobberFinderImpl() BobberFinderImpl {
	return BobberFinderImpl{}
}

type BobberFinderImpl struct {
}

func (bf BobberFinderImpl) FindBobber(knownPosition *image.Point) *image.Point {

	screen := robotgo.CBitmap(robotgo.CaptureScreen())
	var found *image.Point

	search := func(x, y int) bool {
		found = bf.check(screen, x, y)
		if nil != found {
			fmt.Printf("bobber @ %v, %v\n", found.X, found.Y)
		}
		return nil != found
	}

	if nil != knownPosition {
		seachNearby(knownPosition.X, knownPosition.Y, 50, 50, 50, 50, search)
	} else {
		w := 2560
		h := 1440
		w, h = robotgo.GetScreenSize()
		seachNearby(200, 200, 0, w-400, 0, h-400, search)
	}
	return found
}

const verbose = false

func (bf BobberFinderImpl) check(screen robotgo.CBitmap, x int, y int) *image.Point {
	if redFeather(screen, x, y) && blueFeather(screen, x, y) && bobber(screen, x, y) && glow(screen, x, y) && hook(screen, x, y) {
		return &image.Point{X: x, Y: y}
	}
	return nil
}

func seachNearby(x, y, before, after, over, under int, call func(x, y int) bool) bool {
	searchWBegin := max(0, x-before)
	searchWEnd := x + after
	searchHBegin := max(0, y-over)
	searchHEnd := y + under

	for searchY := searchHBegin; searchY < searchHEnd; searchY++ {
		for searchX := searchWBegin; searchX < searchWEnd; searchX++ {
			if call(searchX, searchY) {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func color(screen robotgo.CBitmap, x int, y int) (r int, g int, b int) {
	color := robotgo.CHex(robotgo.GetColor(robotgo.ToMMBitmapRef(screen), x, y))
	return int((color >> 16) & 0xFF), int((color >> 8) & 0xFF), int((color) & 0xFF)
}

func isRed(screen robotgo.CBitmap, x, y int) bool {
	red := 2.0
	minRed := 80
	r, g, b := color(screen, x, y)
	return r >= minRed && r > int(float64(g)*red) && r > int(float64(b)*red)
}

func redFeather(screen robotgo.CBitmap, x, y int) bool {
	if isRed(screen, x, y) {
		redPixelNeeded := 5
		redPixelFound := 0

		if seachNearby(x, y, 2, 2, 2, 2, func(x int, y int) bool {
			if isRed(screen, x, y) {
				redPixelFound++
				if redPixelFound >= redPixelNeeded {
					return true
				}
			}
			return false
		}) {
			if verbose {
				fmt.Printf("maybe bobber RED FEATHER @ %v, %v\n", x, y)
			}
			return true
		}
	}
	return false
}

func blueFeather(screen robotgo.CBitmap, x, y int) bool {

	isBlue := func(screen robotgo.CBitmap, a, c int) bool {
		blue := 1.1
		minBlue := 60
		r, g, b := color(screen, a, c)
		return b >= minBlue && b > int(float64(r)*blue) && b > int(float64(g)*blue) && g > r
	}

	bluePixelNeeded := 5
	bluePixelFound := 0

	if seachNearby(x, y, 5, 5, 20, 0, func(x int, y int) bool {
		if isBlue(screen, x, y) {
			bluePixelFound++
			if bluePixelFound >= bluePixelNeeded {
				return true
			}
		}
		return false
	}) {
		if verbose {
			fmt.Printf("maybe bobber BLUE FEATHER @ %v, %v\n", x, y)
		}
		return true
	}

	return false
}

func bobber(screen robotgo.CBitmap, x, y int) bool {

	isBobber := func(screen robotgo.CBitmap, a, c int) bool {
		bobber := 1.35
		minBlue := 45
		r, g, b := color(screen, a, c)
		return b >= minBlue && r > int(float64(b)*bobber) && g > int(float64(b)*bobber)
	}

	bobberPixel := 5
	bobberPixelFound := 0

	if seachNearby(x, y, 0, 15, 0, 15, func(x int, y int) bool {
		if isBobber(screen, x, y) {
			bobberPixelFound++
			if bobberPixelFound >= bobberPixel {
				return true
			}
		}
		return false
	}) {
		if verbose {
			fmt.Printf("maybe bobber BODY @ %v, %v\n", x, y)
		}
		return true
	}

	return false
}

func glow(screen robotgo.CBitmap, x, y int) bool {

	isGlow := func(screen robotgo.CBitmap, a, c int) bool {
		glow := 1.15
		minBlue := 55
		r, g, b := color(screen, a, c)
		return b >= minBlue && r > int(float64(b)*glow) && g > int(float64(b)*glow)
	}

	glowPixel := 5
	glowPixelFound := 0

	if seachNearby(x, y, 10, 10, 10, 10, func(x int, y int) bool {
		if isGlow(screen, x, y) {
			glowPixelFound++
			if glowPixelFound >= glowPixel {
				return true
			}
		}
		return false
	}) {
		if verbose {
			fmt.Printf("maybe bobber GLOW @ %v, %v\n", x, y)
		}
		return true
	}

	return false
}

func hook(screen robotgo.CBitmap, x, y int) bool {

	isHook := func(screen robotgo.CBitmap, a, c int) bool {
		minR := 70
		r, g, b := color(screen, a, c)
		return r >= minR && r <= g && r <= b
	}

	hookPixel := 5
	hookPixelFound := 0

	if seachNearby(x, y, 0, 15, 15, 0, func(x int, y int) bool {
		if isHook(screen, x, y) {
			hookPixelFound++
			if hookPixelFound >= hookPixel {
				return true
			}
		}
		return false
	}) {
		if verbose {
			fmt.Printf("maybe bobber HOOK @ %v, %v\n", x, y)
		}
		return true
	}

	return false
}

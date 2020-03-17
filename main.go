package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

const (

	xScreenLenght int32 = 800
	yScreenLenght int32 = 600
)

func main() {
	// Initialize sld library
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing sld libs:", err)
		return
	}
	defer sdl.Quit()
	
	// Create a window object, literally a desktop window
	window, err := sdl.CreateWindow("mainWindow", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, xScreenLenght, yScreenLenght, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Creating main game window:", err)
		return
	}
	defer window.Destroy()

	// Create the window background surface
	surface, err := window.GetSurface()
	if err != nil {
		fmt.Println("Creating window surface:", err)
		return
	}
	rect := sdl.Rect{X:0, Y:0, W:xScreenLenght, H:yScreenLenght}
	surface.FillRect(&rect, 0xffffffff)
	window.UpdateSurface()

	// Initialize a rendered object
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing 3d enabled renderer:", err)
		return
	}
	defer renderer.Destroy()

	boyPlayerImg := "sprites/boy_bmps/Idle_1.bmp"
	boyPlayer, boyImgSize, err := newPlayer(renderer, boyPlayerImg)
	if err != nil {
		fmt.Println("Initializing boy player object:", err)
		return
	}

	// Runs until the end of game
	for {
		if running := controlPlayerLoop(); !running {
			return
		}
		renderer.Clear()
		boyPlayer.draw(renderer, boyImgSize)
		renderer.Present()
	}
}
package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

// Screen size
const (
	XScreenLenght int32 = 1800
	YScreenLenght int32 = 800
)

func main() {
	// Initialize sld library
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing sld libs:", err)
		return
	}
	defer sdl.Quit()
	
	// Create a window object, literally a desktop window
	window, err := sdl.CreateWindow("mainWindow", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, XScreenLenght, YScreenLenght, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Creating main game window:", err)
		return
	}
	defer window.Destroy()

	// Initialize a rendered object
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing 3d enabled renderer:", err)
		return
	}
	defer renderer.Destroy()

	boyPlayerImg := "sprites/boy_bmps/small_Idle_1.bmp"
	boyPlayer, err := newPlayer(renderer, boyPlayerImg)
	if err != nil {
		fmt.Println("Initializing boy player object:", err)
		return
	}

	// Runs until the end of game
	for {
		event := sdl.PollEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			fmt.Println("Quit")
			return
		}
		renderer.SetDrawColor(255, 255, 255, 255)
        renderer.Clear()
		boyPlayer.draw(renderer)
		boyPlayer.update()
		renderer.Present()
	}
}
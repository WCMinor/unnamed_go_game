package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

// Screen size
const (
	XScreenLength float64 = 1800
	YScreenLength float64 = 800
)
const (
	spritesPath = "sprites/"
)

// Physics constants
const (
	Gravity float64 = 1
)
// Create an slice with all the elements
var gameElements []*element

func main() {
	// Initialize sld library
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing sld libs:", err)
		return
	}
	defer sdl.Quit()
	
	// Create a window object, literally a desktop window
	window, err := sdl.CreateWindow("mainWindow", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(XScreenLength), int32(YScreenLength), sdl.WINDOW_OPENGL)
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

	gameElements = append(gameElements, newPlayer(renderer, "boy"))
	gameElements = append(gameElements, newStaticPlayer(renderer, "boy"))
	
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
		checkCollisions()
		for _, elm := range gameElements {
			if elm.active {
				err := elm.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
					return
				}
				err = elm.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
			}
		}
		renderer.Present()
	}
}

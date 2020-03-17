package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func controlPlayerLoop() (running bool) {
	running = true
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			running = false
			return running
		}
	}
	return running
}

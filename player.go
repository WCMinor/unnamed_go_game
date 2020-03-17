package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

const (
	xVelocity = 0.1
)
type player struct {
	tex *sdl.Texture
	x, y float32
	H, W int32
}

// Creates a new player rendered texture ready to use
func newPlayer(rendered *sdl.Renderer, spritePath string) (p player, err error) {
	img, err := sdl.LoadBMP(spritePath)
	if err != nil {
		return player{}, fmt.Errorf("Loading player sprite %v", err)
	}
	defer img.Free()
	// Define Height and Weight from the image itself
	p.H = img.H
	p.W = img.W
	// Set starting position of the player
	p.x = float32(XScreenLenght/2)
	p.y = float32(YScreenLenght - p.H)
	// Set texture in the player object
	p.tex, err = rendered.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("Rendering player texture %v", err)
	}
	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	// x and y coordinates at the center of the sprite
	x := int32(p.x) - p.W/2
	renderer.Copy(p.tex,
		&sdl.Rect{X:0, Y:0, H:p.H, W:p.W},
		&sdl.Rect{X:x, Y:int32(p.y), H:p.H, W:p.W})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= xVelocity
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += xVelocity
	}

}


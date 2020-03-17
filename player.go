package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type playerTexture struct {
	tex *sdl.Texture
}

type imgSize struct {
	H int32
	W int32
}

// Creates a new player rendered texture ready to use
func newPlayer(rendered *sdl.Renderer, spritePath string) (player playerTexture, imgSizes imgSize, err error) {
	img, err := sdl.LoadBMP(spritePath)
	if err != nil {
		return playerTexture{}, imgSize{}, fmt.Errorf("Loading player sprite %v", err)
	}
	defer img.Free()
	imgSizes = imgSize{H: img.H, W: img.W}
	player.tex, err = rendered.CreateTextureFromSurface(img)
	if err != nil {
		return playerTexture{}, imgSize{}, fmt.Errorf("Rendering player texture %v", err)
	}
	return player, imgSizes, nil
}

func (player *playerTexture) draw(renderer *sdl.Renderer, imgSizes imgSize) {
	renderer.Copy(player.tex,
		&sdl.Rect{X:0, Y:0, H:imgSizes.H, W:imgSizes.W},
		&sdl.Rect{X:0, Y:0, H:imgSizes.H, W:imgSizes.W})
}
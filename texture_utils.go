package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)


func drawTexture(tex *sdl.Texture, position vector, rotation float64, flip sdl.RendererFlip, renderer *sdl.Renderer) error {
	_, _, height, width, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}
	x := position.x - (float64(width) / 2.0)
	y := position.y - (float64(height) /2.0)

	err = renderer.CopyEx(
		tex,
		&sdl.Rect{X:0, Y:0, H:int32(height), W:int32(width)},
		&sdl.Rect{X:int32(x), Y:int32(y), H:int32(height), W:int32(width)},
		rotation,
		&sdl.Point{Y:(int32(height) / 2), X:(int32(width) / 2)},
		flip)

	if err != nil {
		return fmt.Errorf("Rendering texture: %v", err)
	}
	
	return nil
}

func loadTextureFromBMP(filename string, renderer *sdl.Renderer) (width int32, height int32, tex *sdl.Texture, err error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("Loading %v, %v", filename, err)
	}
	defer img.Free()
	// Set texture
	tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("Rendering texture %v", err)
	}
	width = img.W
	height = img.H

	return width, height, tex, nil
}
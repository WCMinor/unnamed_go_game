package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"strconv"
)

type spriteRenderer struct {
	container *element
	tex *sdl.Texture
	width, height float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer) *spriteRenderer {
	tex := texFromBMP(renderer, spritesPath + container.name + "/Idle_" + strconv.Itoa(container.spritePos) + ".bmp")
	_, _, height, width, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	defer tex.Destroy()
	return &spriteRenderer{
		container: container,
		tex: tex,
		width: float64(width),
		height: float64(height),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	cont := sr.container
	sr.tex.Destroy()
	sr.tex = texFromBMP(renderer, spritesPath + cont.name + "/" + cont.action + "_" + strconv.Itoa(cont.spritePos) + ".bmp")
	x := sr.container.position.x - (sr.width / 2.0)
	y := sr.container.position.y - (sr.height /2.0)

	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X:0, Y:0, H:int32(sr.height), W:int32(sr.width)},
		&sdl.Rect{X:int32(x), Y:int32(y), H:int32(sr.height), W:int32(sr.width)},
		sr.container.rotation,
		&sdl.Point{Y:(int32(sr.height) / 2), X:(int32(sr.width) / 2)},
		sr.container.flip)
	
	return nil
}

func texFromBMP(renderer *sdl.Renderer, filepath string) *sdl.Texture {
	img, err := sdl.LoadBMP(filepath)
	if err != nil {
		panic(fmt.Errorf("Loading %v, %v", filepath, err))
	}
	defer img.Free()
	// Set texture
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Rendering texture %v", err))
	}
	return tex
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}
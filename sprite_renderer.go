package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"path"
)

type spriteRenderer struct {
	container *element
	tex *sdl.Texture
	width, height float64
}

func newSpriteRenderer(container *element, filename string, renderer *sdl.Renderer) *spriteRenderer {
	sr := &spriteRenderer{}
	sr.container = container
	var err error
	_, _, sr.tex, err = loadTextureFromBMP(path.Join(spritesPath, filename), renderer)
	if err != nil {
		panic(fmt.Errorf("loading texture: %v", err))
	}
	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	sr.width = float64(width)
	sr.height = float64(height)
	return sr
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	cont := sr.container
	err := drawTexture(sr.tex, cont.position, cont.rotation, cont.flip, renderer)
	if err != nil {
		panic(fmt.Errorf("drawing texture %v", err))
	}
	
	return nil
}


func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}
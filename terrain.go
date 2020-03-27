package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

func newFloor(renderer *sdl.Renderer, name string) *element {
	floor := &element{}
	floor.active = true
	floor.name = name
	floor.position.x = XScreenLength / 2.0

	spriteRenderer := newSpriteRenderer(floor, "terrain/floor.bmp", renderer)
	fmt.Println(spriteRenderer.height)
	floor.position.y = YScreenLength - spriteRenderer.height / 2.0
	floor.addComponent(spriteRenderer)

	return floor
}
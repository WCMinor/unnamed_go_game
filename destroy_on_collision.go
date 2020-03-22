package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type destroyOnCollision struct {
	container *element
}

func newDestroyOnCollision(container *element) *destroyOnCollision {
	return &destroyOnCollision{
		container: container,
	}
}

func (destoc *destroyOnCollision) onUpdate() error {
	return nil
}

func (destoc *destroyOnCollision) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (destoc *destroyOnCollision) onCollision(other *element) error {
	destoc.container.active = false
	return nil
}
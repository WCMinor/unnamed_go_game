package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type dieOnCollision struct {
	container *element
}

func newDieOnCollision(container *element) *dieOnCollision {
	return &dieOnCollision{
		container: container,
	}
}

func (destoc *dieOnCollision) onUpdate() error {
	if destoc.container.action == "dead" && time.Since(destoc.container.lastMove) > time.Duration(time.Second * 2) {
		destoc.container.active = false
	}
	return nil
}

func (destoc *dieOnCollision) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (destoc *dieOnCollision) onCollision(other *element) error {
	cont := destoc.container
	if cont.action != "dead" {
		cont.action = "dead"
		cont.lastMove = time.Now()
	}

	return nil
}
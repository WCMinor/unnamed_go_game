package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type spritePosUpdater struct {
	container *element
	speed time.Duration
}

func newSpritePosUpdater (container *element) *spritePosUpdater {
	return &spritePosUpdater{
		container: container,
		speed: container.spritePosSpeed,
	}
}

func (pos *spritePosUpdater) onUpdate() error {
	cont := pos.container
	if time.Since(cont.lastSpritePos) > pos.speed && cont.spritePos < 15 {
		cont.spritePos ++
		cont.lastSpritePos = time.Now()
	} else if time.Since(cont.lastSpritePos) > pos.speed && cont.spritePos >= 15 {
		cont.spritePos = 1
		cont.lastSpritePos = time.Now()
	}
	return nil
}

func (pos *spritePosUpdater) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (pos *spritePosUpdater) onCollision(other *element) error {
	return nil
}
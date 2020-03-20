package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type keyboardMover struct {
	container *element
	speed float64
	sr *spriteRenderer
}

func newKeyboardMover (container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed: speed,
		sr: container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	cont := mover.container
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		cont.position.x -= mover.speed
		cont.flip = sdl.FLIP_HORIZONTAL
		cont.action = "Walk"
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		cont.position.x += mover.speed
		cont.flip = sdl.FLIP_NONE
		cont.action = "Walk"
	}
	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardJumper struct {
	container *element
	speed float64
	sr *spriteRenderer
}

func newKeyboardJumper (container *element, speed float64) *keyboardJumper {
	return &keyboardJumper{
		container: container,
		speed: speed,
		sr: container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (jumper *keyboardJumper) onUpdate() error {
	cont := jumper.container
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 {
		cont.position.y -= jumper.speed
	}
	return nil
}

func (jumper *keyboardJumper) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type spritePosUpdater struct {
	container *element
	speed time.Duration
	sr *spriteRenderer
}

func newSpritePosUpdater (container *element, speed time.Duration) *spritePosUpdater {
	return &spritePosUpdater{
		container: container,
		speed: speed,
		sr: container.getComponent(&spriteRenderer{}).(*spriteRenderer),
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

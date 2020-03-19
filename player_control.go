package main

import (
	"github.com/veandco/go-sdl2/sdl"
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
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		cont.position.x += mover.speed
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


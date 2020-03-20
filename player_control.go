package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type keyboardMover struct {
	container *element
	speed float64
}

func newKeyboardMover (container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed: speed,
	}
}

func (mover *keyboardMover) onUpdate() error {
	cont := mover.container
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if ! cont.onLeftWall {
				cont.position.x -= mover.speed
		}
		cont.flip = sdl.FLIP_HORIZONTAL
		cont.action = "Walk"
		cont.lastMove = time.Now()
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if ! cont.onRightWall {
				cont.position.x += mover.speed
		}
		cont.flip = sdl.FLIP_NONE
		cont.action = "Walk"
		cont.lastMove = time.Now()
	}
	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardJumper struct {
	container *element
	speed float64
}

func newKeyboardJumper (container *element, speed float64) *keyboardJumper {
	return &keyboardJumper{
		container: container,
		speed: speed,
	}
}

func (jumper *keyboardJumper) onUpdate() error {
	cont := jumper.container
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if ! cont.onCeiling {
			cont.position.y -= jumper.speed
			cont.action = "Jump"
			cont.lastMove = time.Now()
		}
	}
	return nil
}

func (jumper *keyboardJumper) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type idleDetector struct {
	container *element
	speed time.Duration
}

func newIdleDetector(container *element, speed time.Duration) *idleDetector {
	return &idleDetector{
		container: container,
		speed: speed,
	}
}

func (idle *idleDetector) onUpdate() error {
	cont := idle.container
	if time.Since(cont.lastMove) > idle.speed {
		cont.action = "Idle"
	}
	return nil
}

func (idle *idleDetector) onDraw(renderer *sdl.Renderer) error {
	return nil
}
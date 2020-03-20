package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type keyboardMover struct {
	container *element
	speed float64
}

func newKeyboardMover (container *element) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed: container.xVelocity,
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
		if cont.onFloor {
			cont.action = "Walk"
		}
		cont.lastMove = time.Now()
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if ! cont.onRightWall {
				cont.position.x += mover.speed
		}
		cont.flip = sdl.FLIP_NONE
		if cont.onFloor {
			cont.action = "Walk"
		}
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

func newKeyboardJumper (container *element) *keyboardJumper {
	return &keyboardJumper{
		container: container,
		speed: container.yVelocity,
	}
}

func (jumper *keyboardJumper) onUpdate() error {
	cont := jumper.container
	if ! cont.onCeiling {
		if time.Since(cont.startJump) < (cont.spritePosSpeed * time.Duration(cont.spritesNum))/2 {
			cont.position.y -= jumper.speed
			cont.action = "Jump"
			cont.lastMove = time.Now()
		} else if time.Since(cont.startJump) < (cont.spritePosSpeed * time.Duration(cont.spritesNum)) {
			cont.action = "Jump"
			cont.lastMove = time.Now()
		}
		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_SPACE] == 1 {
			if cont.onFloor {
				cont.startJump = time.Now()
			}
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

func newIdleDetector(container *element) *idleDetector {
	return &idleDetector{
		container: container,
		speed: container.moveSpeed,
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
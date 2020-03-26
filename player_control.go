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
				cont.position.x -= mover.speed * delta
				for i := range cont.collisionPoints {
					cont.collisionPoints[i].center = cont.position
				}
		}
		cont.flip = sdl.FLIP_HORIZONTAL
		if cont.onFloor {
			cont.action = "walk"
		}
		cont.lastMove = time.Now()
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if ! cont.onRightWall {
				cont.position.x += mover.speed * delta
				for i := range cont.collisionPoints {
					cont.collisionPoints[i].center = cont.position
				}
		}
		cont.flip = sdl.FLIP_NONE
		if cont.onFloor {
			cont.action = "walk"
		}
		cont.lastMove = time.Now()
	}
	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
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
	if ! cont.onCeiling && cont.jumping {
		if (YScreenLength - cont.position.y) <= cont.jumpHigh {
			cont.position.y -= Gravity * delta * cont.yVelocity
			for i := range cont.collisionPoints {
				cont.collisionPoints[i].center = cont.position
			}
			cont.action = "jump"
			cont.lastMove = time.Now()
		} else {
			cont.jumping = false
		}
	}
	if cont.onFloor {
		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_SPACE] == 1 {
			cont.jumping = true
		}
	}
	return nil
}

func (jumper *keyboardJumper) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (jumper *keyboardJumper) onCollision(other *element) error {
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
	if time.Since(cont.lastMove) > idle.speed && cont.action != "dead" {
		cont.action = "idle"
	}
	return nil
}

func (idle *idleDetector) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (idle *idleDetector) onCollision(other *element) error {
	return nil
}
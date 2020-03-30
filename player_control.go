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
	on := cont.getComponent(&onSurface{}).(*onSurface)
	jumper := cont.getComponent(&jumper{}).(*jumper)
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		moveLeft(cont, on, mover.speed)
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		moveRight(cont, on, mover.speed)
	}
	if on.H && keys[sdl.SCANCODE_SPACE] == 1 {
		jumper.jumping = true
		jumper.jumpStartHeight = cont.position.y
		on.H = false
	}
	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

type jumper struct {
	container *element
	jumpSpeed, jumpHeight, jumpStartHeight float64
	jumping bool
}

func newJumper (container *element, jumpSpeed, jumpHeight float64) *jumper {
	return &jumper{
		container: container,
		jumpSpeed: jumpSpeed,
		jumpHeight: jumpHeight,
	}
}

func (jumper *jumper) onUpdate() error {
	cont := jumper.container
	if jumper.jumping {
		if (jumper.jumpStartHeight - cont.position.y) <= jumper.jumpHeight {
			cont.position.y -= Gravity * delta * jumper.jumpSpeed
			moveCollisions(cont)
			cont.action = "jump"
			cont.lastMove = time.Now()
			
		} else {
			jumper.jumping = false
		}
	}
	return nil
}

func (jumper *jumper) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (jumper *jumper) onCollision(other *element) error {
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

func moveCollisions(container *element) {
	for i := range container.collisionPoints {
		container.collisionPoints[i].center = container.position
	}
	for i := range container.collisionRects {
		container.collisionRects[i].center = container.position
	}
}
func moveDown(container *element, speed float64) {
	container.position.y += speed * delta
	moveCollisions(container)
}

func moveLeft(container *element, on *onSurface, speed float64) {
	if ! on.V {
			container.position.x -= speed * delta
			moveCollisions(container)
	}
	container.flip = sdl.FLIP_HORIZONTAL
	if on.H {
		container.action = "walk"
	}
	container.lastMove = time.Now()
}

func moveRight(container *element, on *onSurface, speed float64) {
	if ! on.V {
			container.position.x += speed * delta
			moveCollisions(container)
	}
	container.flip = sdl.FLIP_NONE
	if on.H {
		container.action = "walk"
	}
	container.lastMove = time.Now()

}

func jump(container *element, on *onSurface) {

}
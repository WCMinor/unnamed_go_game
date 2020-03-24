package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gravity struct {
	container *element
	speed float64
}

func newGravity (container *element) *gravity {
	return &gravity{
		container: container,
		speed: Gravity,
	}
}

func (g *gravity) onUpdate() error {
	cont := g.container
	if ! cont.onFloor {
		cont.position.y += g.speed * delta
		for i := range cont.collisionPoints {
			cont.collisionPoints[i].center = cont.position
		}
	}
	return nil
}

func (g *gravity) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (g *gravity) onCollision(other *element) error {
	return nil
}

type onSurface struct {
	container *element
}

func newOnSurface (container *element) *onSurface {
	return &onSurface{
		container: container,
	}
}

func (ons *onSurface) onUpdate() error {
	cont := ons.container
	if (cont.position.y + cont.height / 2) >= YScreenLength {
		cont.onFloor = true
	} else {
		cont.onFloor = false
	}
	if (cont.position.y - cont.height / 2) <= 0 {
		cont.onCeiling = true
	} else {
		cont.onCeiling = false
	}
	if (cont.position.x - cont.width /2 ) <= 0 {
		cont.onLeftWall = true
	} else if (cont.position.x + cont.width /2 ) >= XScreenLength {
		cont.onRightWall = true
	} else {
		cont.onRightWall = false
		cont.onLeftWall = false
	}
	return nil
}

func (ons *onSurface) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (ons *onSurface) onCollision(other *element) error {
	return nil
}
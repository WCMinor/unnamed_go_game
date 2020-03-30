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
	on := cont.getComponent(&onSurface{}).(*onSurface)

	if ! on.H {
		moveDown(cont, g.speed)
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
	animator * animator
	H, V bool
}

func newOnSurface (container *element) *onSurface {
	return &onSurface{
		container: container,
		animator: container.getComponent(&animator{}).(*animator),
		H: false,
		V: false,
	}
}

func (ons *onSurface) onUpdate() error {
	cont := ons.container
	height := ons.animator.height
	width := ons.animator.width
	if (cont.position.y + height / 2) >= YScreenLength {
		//ons.H = true
	} else {
		//ons.H = false
	}
	if (cont.position.x - width /2 ) <= 0 || (cont.position.x + width /2 ) >= XScreenLength {
		//ons.V= true
	} else {
		//ons.V = false
	}
	return nil
}

func (ons *onSurface) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (ons *onSurface) onCollision(other *element) error {
	ons.H = true
	return nil
}
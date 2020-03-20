package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	xVelocity = 0.5
	yVelocity = 5
	spritesNum int32 = 15
	posSpeed time.Duration = time.Millisecond * 40 //milliseconds
	moveSpeed time.Duration = time.Millisecond * 160 //milliseconds
)


func newPlayer(renderer *sdl.Renderer, name string) *element {
	player := &element{}
	player.active = true
	player.action = "Idle"
	player.name = name
	player.spritePos = 1
	player.position.x = XScreenLength / 2.0

	sr := newSpriteRenderer(player, renderer)
	player.addComponent(sr)

	player.height = sr.height
	player.width = sr.width
	player.position.y = YScreenLength - player.height

	gravity := newGravity(player, Gravity)
	player.addComponent(gravity)
	mover := newKeyboardMover(player, xVelocity)
	player.addComponent(mover)
	jumper := newKeyboardJumper(player, yVelocity)
	player.addComponent(jumper)
	sPosUpdater := newSpritePosUpdater(player, posSpeed)
	player.addComponent(sPosUpdater)
	ons := newOnSurface(player)
	player.addComponent(ons)
	idle := newIdleDetector(player, moveSpeed)
	player.addComponent(idle)


	return player
}
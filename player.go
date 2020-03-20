package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func newPlayer(renderer *sdl.Renderer, name string) *element {
	player := &element{}
	player.active = true
	player.action = "Idle"
	player.name = name
	player.spritePos = 1
	player.position.x = XScreenLength / 2.0
	player.xVelocity = 0.5
	player.yVelocity = 2
	player.spritesNum = 15
	player.spritePosSpeed = time.Millisecond * 40 //milliseconds
	player.moveSpeed = time.Millisecond * 160 //milliseconds


	sr := newSpriteRenderer(player, renderer)
	player.addComponent(sr)

	player.height = sr.height
	player.width = sr.width
	player.position.y = YScreenLength - player.height

	gravity := newGravity(player)
	player.addComponent(gravity)
	mover := newKeyboardMover(player)
	player.addComponent(mover)
	jumper := newKeyboardJumper(player)
	player.addComponent(jumper)
	sPosUpdater := newSpritePosUpdater(player)
	player.addComponent(sPosUpdater)
	ons := newOnSurface(player)
	player.addComponent(ons)
	idle := newIdleDetector(player)
	player.addComponent(idle)

	return player
}
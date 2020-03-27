package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"path"
	"fmt"
)

func newPlayer(renderer *sdl.Renderer, name string) *element {
	var sampleRate float64 = 20
	player := &element{}
	player.active = true
	player.action = "idle"
	player.name = name
	player.position.x = XScreenLength / 2.0
	player.xVelocity = 10
	player.yVelocity = 2
	player.jumpHeight = 250
	player.moveSpeed = time.Millisecond * 160 //milliseconds

	sequences := make(map[string]*sequence)

	sequenceMap :=map[string]bool{
		"idle": true,
		"walk": true,
		"run": true,
		"jump": false,
		"dead": false,
	}

	for seq, loop := range sequenceMap {
		sequence, err := newSequence(path.Join(spritesPath, player.name, seq), sampleRate, loop, renderer)
		if err != nil {
			panic(fmt.Errorf("loading textures sequence: %v", err))
		}
		sequences[seq] = sequence
	}
	animator, err := newAnimator(player, sequences)
	if err != nil {
		panic(fmt.Errorf("Creating new animator: %v", err))
	}
	player.addComponent(animator)

	player.position.y = YScreenLength - animator.width

	gravity := newGravity(player)
	player.addComponent(gravity)
	mover := newKeyboardMover(player)
	player.addComponent(mover)
	jumper := newKeyboardJumper(player)
	player.addComponent(jumper)
	ons := newOnSurface(player)
	player.addComponent(ons)
	idle := newIdleDetector(player)
	player.addComponent(idle)

	colPoint := circle{
		center: player.position,
		radius: animator.width / 3,
	}

	player.collisionPoints = append(player.collisionPoints, colPoint)

	return player
}
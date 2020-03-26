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
	player.jumpHigh = 250
	player.moveSpeed = time.Millisecond * 160 //milliseconds


	sequences := make(map[string]*sequence)

	sequenceList :=[]string{
		"idle",
		"walk",
		"run",
		"jump",
		"dead",
	}

	for _, seq := range sequenceList {
		sequence, err := newSequence(path.Join(spritesPath, player.name, seq), sampleRate, true, renderer)
		if err != nil {
			panic(fmt.Errorf("loading textures sequence: %v", err))
		}
		sequences[seq] = sequence
	}
	animator := newAnimator(player, sequences)
	player.addComponent(animator)
	width, height, _, err := loadTextureFromBMP(path.Join(spritesPath, player.name, "idle/01.bmp"), renderer)
	if err != nil {
		panic(fmt.Errorf("getting info from bmp: %v", err))
	}
	player.width = float64(width)
	player.height = float64(height)
	player.position.y = YScreenLength - player.height

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
		radius: player.width / 3,
	}

	player.collisionPoints = append(player.collisionPoints, colPoint)

	return player
}

func newStaticPlayer(renderer *sdl.Renderer, name string) *element {
	var sampleRate float64 = 20
	player := &element{}
	player.active = true
	player.action = "idle"
	player.name = name
	player.position.x = XScreenLength / 1.5
	player.xVelocity = 0.5 * delta
	player.yVelocity = 10.5 * delta
	player.moveSpeed = time.Millisecond * 160 //milliseconds

	sequences := make(map[string]*sequence)

	sequencesMap := map[string]bool{
		"idle": true,
		"walk": true,
		"run": true,
		"jump": true,
		"dead": false,
	}

	for seq, loop := range sequencesMap {
		sequence, err := newSequence(path.Join(spritesPath, player.name, seq), sampleRate, loop, renderer)
		if err != nil {
			panic(fmt.Errorf("loading textures sequence: %v", err))
		}
		sequences[seq] = sequence
	}
	animator := newAnimator(player, sequences)
	player.addComponent(animator)
	width, height, _, err := loadTextureFromBMP(path.Join(spritesPath, player.name, "idle/01.bmp"), renderer)
	if err != nil {
		panic(fmt.Errorf("getting info from bmp: %v", err))
	}
	player.width = float64(width)
	player.height = float64(height)
	player.position.y = YScreenLength - player.height

	gravity := newGravity(player)
	player.addComponent(gravity)
	ons := newOnSurface(player)
	player.addComponent(ons)
	idle := newIdleDetector(player)
	player.addComponent(idle)
	colDestroy := newDieOnCollision(player)
	player.addComponent(colDestroy)

	colPoint := circle{
		center: player.position,
		radius: player.width / 3,
	}

	player.collisionPoints = append(player.collisionPoints, colPoint)

	return player
}
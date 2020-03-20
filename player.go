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

/*
// Creates a new player rendered texture ready to use
func newPlayer(renderer *sdl.Renderer, name string) (p player) {
	p.H = playerHeight
	p.W = playerWidth
	p.name = name
	p.actpos = "1"
	p.jumpPos = 1
	p.action = "Idle"
	p.sense = "right"
	p = loadFromBMP(p, renderer, spritesPath + p.name + "/" + p.action + "_" + p.actpos + ".bmp")

	// Set starting position of the player
	p.x = float32(XScreenLength/2)
	p.y = float32(YScreenLength - p.H)
	// Set texture in the player object
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	// load a new texture each renderization
	p.tex = loadFromBMP(*p, renderer, spritesPath + p.name + "/" + p.action + "_" + p.actpos + ".bmp").tex
	// x and y coordinates at the center of the sprite
	x := int32(p.x) - p.W / 2
	flip := sdl.FLIP_NONE
	if p.sense == "left" {
		flip = sdl.FLIP_HORIZONTAL
	}

	renderer.CopyEx(p.tex,
		&sdl.Rect{X:0, Y:0, H:p.H, W:p.W},
		&sdl.Rect{X:x, Y:int32(p.y), H:p.H, W:p.W},
		0.0,
		&sdl.Point{Y: p.H / 2, X: p.W / 10},
		flip)
}

func (p *player) incActpos() {
	if p.walking || p.idle {

	}
}

func (p *player) onFloor() bool {
	if p.y < float32(YScreenLength-p.H) {
		return false
	}
	return true
}
func (p *player) gravity() {
	if !p.onFloor() && ! p.jumping {
		p.y += Gravity
	}
}

func (p *player) jump() {
// The player should be jumping during exactly 15 positions to match the 15 sprites	
	if p.jumping {
		p.idle = false
		p.action ="Jump"
		if p.jumpPos < spritesNum {
			if time.Since(p.lastMove) > moveSpeed {
				p.y -= yVelocity
				p.actpos = strconv.Itoa(int(p.jumpPos))
				p.jumpPos ++
				p.lastMove = time.Now()
			}
		} else {
			p.jumping = false
			p.jumpPos = 1
		}
	}
}

func (p *player) move() {
	p.lastMove = time.Now()
	p.idle = false
	if p.sense == "left" {
		p.x -= xVelocity
	} else if p.sense == "right" {
		p.x += xVelocity
	}

	if p.onFloor() {
		p.walking = true
		p.action = "Walk"
	}
}

func (p *player) stayIdle() {
	if time.Since(p.lastMove) > moveSpeed {
		p.idle = true
		p.action = "Idle"
	}
}

func (p *player) update() {
	// Gravity is inevitable
	p.gravity()
	p.jump()
	p.incActpos()
	p.stayIdle()
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 && p.onFloor() {
		p.jumping = true
	} else if keys[sdl.SCANCODE_LEFT] == 1 {
		p.sense = "left"
		p.move()
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.sense = "right"
		p.move()
	}
}
*/
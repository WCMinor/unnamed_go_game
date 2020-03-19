package main

import (
	"github.com/veandco/go-sdl2/sdl"
//	"fmt"
	"time"
//	"strconv"
)

const (
	xVelocity = 0.5
	yVelocity = 1
	spritesNum int32 = 15
	spritesPath = "sprites/"
	playerHeight float64 = 200
	playerWidth float64 = 184
	posSpeed = time.Millisecond * 40 //milliseconds
	moveSpeed = time.Millisecond * 60 //milliseconds
)
/*
type player struct {
	tex *sdl.Texture
	name, action, actpos, sense string
	x, y float32
	H, W, jumpPos int32
	lastMove, lastPos time.Time
	walking, jumping, idle bool
}
*/


func newPlayer(renderer *sdl.Renderer, name string) *element {
	player := &element{}
	player.active = true
	player.name = name
	player.position.x = XScreenLength / 2.0
	player.position.y = YScreenLength - playerHeight

	sr := newSpriteRenderer(player, renderer, spritesPath + player.name + "/Idle_1.bmp")
	player.addComponent(sr)

	mover := newKeyboardMover(player, xVelocity)
	player.addComponent(mover)
	jumper := newKeyboardJumper(player, yVelocity)
	player.addComponent(jumper)
	
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
		if time.Since(p.lastPos) > posSpeed {
			pos, err := strconv.Atoi(p.actpos)
			if err != nil {
				panic(fmt.Errorf("converting action position to integer %v", err))
			}
			if pos < 15 {
				pos ++
			} else {
				pos = 1
			}
			p.actpos = strconv.Itoa(pos)
			p.lastPos = time.Now()
		}
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
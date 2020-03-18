package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"time"
	"strconv"
)

const (
	xVelocity = 0.1
	spritesPath = "sprites/"
	stepsSpeed = time.Millisecond * 20 //milliseconds
)
type player struct {
	tex *sdl.Texture
	name, action, actpos, sense string
	x, y float32
	H, W int32
	lastStep time.Time
}

func loadFromBMP(p player, renderer *sdl.Renderer, filepath string) player {
	img, err := sdl.LoadBMP(filepath)
	if err != nil {
		panic(fmt.Errorf("Loading %v, %v", filepath, err))
	}
	defer img.Free()
	// Set texture
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Rendering texture %v", err))
	}
	
	// Define Height and Weight from the image itself
	p.H = img.H
	p.W = img.W

	return p
}

// Creates a new player rendered texture ready to use
func newPlayer(renderer *sdl.Renderer, name string) (p player) {
	p.name = name
	p.actpos = "1"
	p.action = "Idle"
	p.sense = "right"
	p = loadFromBMP(p, renderer, spritesPath + p.name + "/" + p.action + "_" + p.actpos + ".bmp")

	// Set starting position of the player
	p.x = float32(XScreenLenght/2)
	p.y = float32(YScreenLenght - p.H)
	// Set texture in the player object
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	// x and y coordinates at the center of the sprite
	p.tex = loadFromBMP(*p, renderer, spritesPath + p.name + "/" + p.action + "_" + p.actpos + ".bmp").tex
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
	if time.Since(p.lastStep) > stepsSpeed {
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
		p.lastStep = time.Now()
	}
}
func (p *player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		// Walking
		p.x -= xVelocity
		p.action = "Walk"
		p.sense = "left"
		p.incActpos()

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += xVelocity
		p.action = "Walk"
		p.sense = "right"
		p.incActpos()
	} else {
		p.action = "Idle"
	}
}
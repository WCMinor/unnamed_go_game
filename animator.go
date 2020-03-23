package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"io/ioutil"
	"fmt"
	"path"
)

type animator struct {
	container *element
	sequences map[string]*sequence
	currentSequence string
	lastFrameChange time.Time
}

func newAnimator(container *element, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator
	an.container = container
	an.sequences = sequences
	an.currentSequence = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) onDraw(renderer *sdl.Renderer) error {
	tex := an.sequences[an.currentSequence].texture()

	return drawTextue(tex, an.container.position, an.container.rotation, an.container.flip, renderer)
}

func (an *animator) onUpdate() error {
	sequence := an.sequences[an.currentSequence]
	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(an.lastFrameChange) > time.Duration(frameInterval) {
		sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}

	return nil
}

func (an *animator) onCollision(other *element) error {
	return nil
}

type sequence struct {
	textures []*sdl.Texture
	frame int
	sampleRate float64
	loop bool
}

func newSequence(filepath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*sequence, error) {
	var seq sequence
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading files from %v: %v", filepath, err)
	}
	for _, file := range(files) {
		filename := path.Join(filepath, file.Name())
		tex, err := loadTextureFromBMP(filename, renderer)
		if err != nil {
			return nil, fmt.Errorf("loading texture from file: %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}
	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil

}

// Returns the currently active texture
func (seq *sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

// Increments the sequence
func (seq *sequence) nextFrame() {
	if seq.frame == len(seq.textures) -1 {
		if seq.loop {
			seq.frame = 1
		}
	} else {
		seq.frame ++
	}
}
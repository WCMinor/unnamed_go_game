package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	"fmt"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type element struct {
	name string
	active bool
	position vector
	rotation float64
	flip sdl.RendererFlip
	components []component
}

func (elem *element) addComponent(new component) {
	// Panic if a new component gets created using an existing type
	for _, existing := range elem.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attemp to add a new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	// Add the new component to the components array
	elem.components = append(elem.components, new)
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	panic(fmt.Sprintf(
		"no component found with type of %v",
		reflect.TypeOf(withType)))
}
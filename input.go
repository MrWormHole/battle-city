package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	ownerEntity *Entity
	speed       float64
}

func createInputComponent(entity *Entity, speed float64) *Input {
	return &Input{ownerEntity: entity,
		speed: speed}
}

func (input Input) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && input.ownerEntity.position.x > 0 {
		input.ownerEntity.position.x -= input.speed * timer.deltaTime
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && input.ownerEntity.position.x < SCREEN_WIDTH {
		input.ownerEntity.position.x += input.speed * timer.deltaTime
	} else if keys[sdl.SCANCODE_UP] == 1 && input.ownerEntity.position.y > 0 {
		input.ownerEntity.position.y -= input.speed * timer.deltaTime
	} else if keys[sdl.SCANCODE_DOWN] == 1 && input.ownerEntity.position.y < SCREEN_HEIGHT {
		input.ownerEntity.position.y += input.speed * timer.deltaTime
	}

	return nil
}
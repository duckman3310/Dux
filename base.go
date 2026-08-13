package dux

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Pos  rl.Vector2
	Size rl.Vector2
}

func (e *Entity) GetEntity() *Entity {
	return e
}

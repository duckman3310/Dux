package elements

import (
	"dux"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Frame - рамка/каркас, отрисовывает простой прямоугольник,
// нужен как основа для элементов по сложнее
type Frame struct {
	dux.Entity
}

// Возвращяет ссылку на созданый Frame
func NewFrame() *Frame {
	return &Frame{
		Segments: 8,
	}

}

func (f *Frame) Draw() {
	rect := rl.NewRectangle(f.Pos.X, f.Pos.Y, f.Size.X, f.Pos.Y)

	if f.Color.A > 0 {
		rl.DrawRectangleRounded(rect, f.Roundness, f.Segments, f.Color)
	}

	if f.BorderColor.A > 0 && f.BorderWidth > 0 {
		rl.DrawRectangleLinesEx(rect, f.BorderWidth, f.BorderColor)
	}
}

func (f *Frame) Update() {

}

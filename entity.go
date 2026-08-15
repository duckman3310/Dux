package dux

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Pos  rl.Vector2 // позиция
	Size rl.Vector2 // размер

	Color           rl.Color // основной цвет
	BackgroundColor rl.Color // цвет фона

	BorderColor rl.Color // цвет бордюра
	BorderWidth float32  // ширина бордюра

	Roundness float32 // скругление, от 0.0 до 1.0
	Segments  int32   // количество итераций скругления

	// внешний отступ
	Margin       float32
	MarginTop    float32
	MarginRight  float32
	MarginLeft   float32
	MarginBottom float32

	// Внутрений отступ
	Padding       float32
	PaddingTop    float32
	PaddingRight  float32
	PaddingLeft   float32
	PaddingBottom float32
}

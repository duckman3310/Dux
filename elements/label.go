package elements

import (
	base "dux"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	base.Entity
	Text string

	Font      rl.Font
	FontSize  float32
	FontColor rl.Color
	Spacing   float32
	AlignX    float32 // 0.0 = left, 0.5 = center, 1.0 = right
	AlignY    float32 // 0.0 = top,  0.5 = center, 1.0 = bottom
}

func NewLabel(text string) *Label {
	return &Label{
		Text: text,
	}
}

func (l *Label) Draw() {

	textSize := rl.MeasureTextEx(l.Font, l.Text, l.FontSize, l.Spacing)

	textPos := rl.Vector2{
		X: l.Pos.X + (l.Size.X-textSize.X)*l.AlignX,
		Y: l.Pos.Y + (l.Size.Y-textSize.Y)*l.AlignY,
	}

	rl.DrawTextEx(l.Font, l.Text, textPos, l.FontSize, l.Spacing, l.FontColor)
}
func (l *Label) Update() {

}

func (l *Label) SetText(text string) {
	l.Text = text
}

func (l *Label) GetText() string {
	return l.Text
}

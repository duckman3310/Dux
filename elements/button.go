package elements

import (
	base "dux"
)

type Button struct {
	base.Entity

	Frame *Frame
	Label *Label

	Actions []func()
}

func NewButton(text string, actions ...func()) *Button {
	return &Button{
		Frame:   NewFrame(),
		Label:   NewLabel(text),
		Actions: actions,
	}
}

func (b *Button) Draw() {
	b.Frame.Draw()
	b.Label.Draw()

}

func (b *Button) Update() {
	b.Frame.Pos = b.Pos
	b.Frame.Size = b.Size

	b.Label.Pos = b.Pos
	b.Label.Size = b.Size
}

func (b *Button) SetText(text string) {
	b.Label.Text = text
}

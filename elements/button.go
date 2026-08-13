package elements

import (
	base "dux"
)

type Button struct {
	base.Entity

	text   string
	action func()
}

func NewButton(text string, action func()) *Button {
	return &Button{
		text:   text,
		action: action,
	}
}

func (b *Button) Draw() {

}

func (b *Button) Update() {

}

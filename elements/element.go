package elements

import (
	base "dux"
)

type Element interface {
	Draw()
	Update()

	GetEntity() *base.Entity
}

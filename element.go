package dux

type Element interface {
	Draw()
	Update()

	GetEntity() *Entity
}

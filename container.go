package dux

type Layout interface {
	Arrange(c *Container)
}

type Container struct {
	Entity
	Elements []*Element
	Layout   Layout
}

func NewContainer() *Container {
	return &Container{
		Elements: make([]*Element, 0),
	}
}

func (c *Container) Draw() {

	// рисуем все элементы контейнера при их наличии
	if len(c.Elements) > 0 {
		for _, element := range c.Elements {
			element.Draw()
		}
	}
}

func (c *Container) Update() {

	// обновляем все элементы контейнера при их наличии
	if len(c.Elements) > 0 {
		for _, element := range c.Elements {
			element.Update()
		}
	}
}

func (c *Container) None() {

}

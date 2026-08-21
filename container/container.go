package container

import (
	"dux"
)

// Layout определяет интерфейс для алгоритмов расстановки элементов внутри контрейнера.
// чтобы его реализовать нужно создать структуру и метод Arrange(c *Container) для нее
type Layout interface {
	Arrange(c *Container)
}

// Container - элемент содержащий другие элементы
// и управляющий их позиционированием
type Container struct {
	dux.Entity
	Elements []dux.Element // дочерние элементы
	Layout   Layout        // алгоритм расстановки элементов
}

// возвращяет ссылку на созданый контейнер
func NewContainer() *Container {
	return &Container{
		Elements: make([]dux.Element, 0),
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

		c.Layout.Arrange(c)

		for _, element := range c.Elements {
			element.Update()
		}
	}
}

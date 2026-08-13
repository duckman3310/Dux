package container

// LinearLayout разтавляет элементы в одну линию
// по горизонтали или вертикали, разтягивая их по пемпедекулярной оси
type LinearLayout struct {
	orientation string
}

// реализация метода для интерфейса Layout
func (f *LinearLayout) Arrange(c *Container) {

	// приводим любую ориентацию к простому флагу "горизонтально или нет"
	isHorizontal := f.orientation == "horizontal"
	if f.orientation == "adaptive" {
		isHorizontal = c.Size.X > c.Size.Y
	}

	// определяем начальную позицию и размер, в зависимости от ориентации
	// lastPos (или main ниже) - позиция и размер относительно основной оси, cross - поперечной
	lastPos := c.Pos.X
	crossPos := c.Pos.Y
	crossSize := c.Size.Y

	if !isHorizontal {
		lastPos = c.Pos.Y
		crossPos = c.Pos.X
		crossSize = c.Size.X
	}

	// алгоритм растановки
	for _, element := range c.Elements {

		// берем ссылку на entity текущего элемента
		entity := element.GetEntity()

		// создаем ссылки на координаты и размеры entity
		// в зависимости от ориентации
		mainP, mainS := &entity.Pos.X, &entity.Size.X
		crossP, crossS := &entity.Pos.Y, &entity.Size.Y

		if !isHorizontal {
			mainP, mainS = &entity.Pos.Y, &entity.Size.Y
			crossP, crossS = &entity.Pos.X, &entity.Size.X
		}

		// задаем позицию и размер
		*mainP = lastPos
		lastPos += *mainS

		*crossP = crossPos
		*crossS = crossSize
	}
}

package animation

type Animation struct {
	Target    *float32
	StartVal  float32 // Откуда начинаем
	TargetVal float32 // Куда идем
	Progress  float32 // От 0.0 до 1.0
	Duration  float32 // Длительность в секундах
	AnimFunc  func(a *Animation, dt float32)
}

type AnimManager struct {
	animations []*Animation
}

func NewAnimManager() *AnimManager {
	return &AnimManager{
		animations: make([]*Animation, 0),
	}
}

func (m *AnimManager) AddAnimation(target *float32, targetVal, duration float32, animFunc func(a *Animation, dt float32)) {
	anim := &Animation{
		Target:    target,
		StartVal:  *target,
		TargetVal: targetVal,
		Progress:  0,
		AnimFunc:  animFunc,
	}
	m.animations = append(m.animations, anim)
}

// Update() выполняет шаг каждой анимации
// из слайса, удаляет если она завершилась
func (m *AnimManager) Update(dt float32) {
	n := 0 // указатель

	// проходимся по каждой (если они есть)
	if len(m.animations) > 0 {
		for _, anim := range m.animations {

			// выполняем шаг анимации
			anim.AnimFunc(anim, dt)

			// если анимация не завершена перемещяем ее на
			// указатель и сдвигаем его дальше по слайсу
			if anim.Progress < 1.0 {
				m.animations[n] = anim
				n++

			} else {
				// если завершина ниче не делаем, на ее место
				// будет записана ссылка на другую или nil

				// выравниваем значение
				*anim.Target = anim.TargetVal
			}
		}

		// чтобы сборщик убрал из памяти завершенные анимации,
		// нам нужно убрать на них ссылки

		// проходимся по всем анимациям начиная с указателя (завершенных)
		// и меняем ссылку на nil
		for i := n; i < len(m.animations); i++ {
			m.animations[i] = nil
		}

		// срезаем у слайса все после указателя
		m.animations = m.animations[:n]
	}
}

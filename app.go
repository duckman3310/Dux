package dux

// App - "настройки" вашего приложения,
// создав его, в последсвии вы можете передавать его как
// опциональный аргумент для New функций,
// они будут брать стандартные значения из этой структуры

import rl "github.com/gen2brain/raylib-go/raylib"

type App struct {

	// шрифт
	Font     rl.Font
	FontSize float32 // размер шрифта
	Spacing  float32 // растояние между буквами
}

func defaultApp() *App {
	return &App{
		Font:     rl.GetFontDefault(),
		FontSize: 20,
		Spacing:  1,
	}
}

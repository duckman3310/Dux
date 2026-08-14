package font

import (
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// LoadFont загружает один шрифт по укзаному пути,
// шрифт содержит полный набор (ASCII + кириллица) с указаным размером
func LoadFont(path string, size int32) rl.Font {

	var codepoints []rune // руна это номер символа в таблице Unicode

	// в два захода наполняем слайс номерами нужных символов,
	// в нашем случае ASCII и кириллица

	for r := rune(32); r <= 126; r++ {
		codepoints = append(codepoints, r)
	}

	for r := rune(0x0400); r <= 0x04FF; r++ {
		codepoints = append(codepoints, r)
	}

	// получаем нужный шрифт в raylib формате и возвращяем
	return rl.LoadFontEx(path, size, codepoints)
}

// LoadFontDir загружает все шрифты из директории по указаному пути,
// шрифт содержит полный набор (ASCII + кириллица) с указаным размером
func LoadFontDir(path string, size int32) ([]rl.Font, error) {

	// получаем слайс "легких" файлов по указаному пути
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var Fonts []rl.Font

	// проходимся по каждому файлу игнорируя папки и не шрифты
	for _, entry := range entries {
		if !entry.IsDir() {
			ext := filepath.Ext(entry.Name())
			if ext == ".ttf" || ext == ".otf" {

				// создаем путь текущего шрифта, и добавляем
				// в наш слайс ранее написаной функциией
				fontPath := filepath.Join(path, entry.Name())
				Fonts = append(Fonts, LoadFont(fontPath, size))
			}
		}
	}
	return Fonts, nil
}

package font

import (
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// коментарии потом добавлю

type FontManager struct {
	Fonts map[string]rl.Font
}

func NewFontManager() *FontManager {
	return &FontManager{
		Fonts: make(map[string]rl.Font),
	}
}
func (fm *FontManager) LoadFont(name string, size int32, path string) {
	fm.Fonts[name] = LoadFont(path, size)

}

func (fm *FontManager) LoadFontDir(path string, size int32) error {
	if fm.Fonts == nil {
		fm.Fonts = make(map[string]rl.Font)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := filepath.Ext(entry.Name())
		if ext == ".ttf" || ext == ".otf" {

			name := strings.TrimSuffix(entry.Name(), ext)
			fontPath := filepath.Join(path, entry.Name())

			fm.Fonts[name] = LoadFont(fontPath, size)
		}
	}

	return nil
}

func (fm *FontManager) Get(name string) rl.Font {
	return fm.Fonts[name]
}

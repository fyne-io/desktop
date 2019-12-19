package desktop

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	wmtheme "fyne.io/desktop/theme"
)

func updateWallpaperPath(background *canvas.Image, path string) {
	if path != "" {
		background.Resource = nil
		background.File = path
		return
	}
	background.Resource = wmtheme.Background
	background.File = ""
}

func wallpaperPath() string {
	pathEnv := Instance().Settings().Background()
	if pathEnv == "" {
		return ""
	}

	if stat, err := os.Stat(pathEnv); os.IsNotExist(err) || !stat.Mode().IsRegular() {
		return ""
	}

	return pathEnv
}

func newBackground() fyne.CanvasObject {
	imagePath := wallpaperPath()
	if imagePath != "" {
		return canvas.NewImageFromFile(imagePath)
	}
	return canvas.NewImageFromResource(wmtheme.Background)
}

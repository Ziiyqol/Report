package app

import (
	"os"
	"report/internal/config"
	"report/internal/services"
	"report/internal/storage"
	"report/internal/ui"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

// Run запускает приложение.
func Run() {
	// Создаём fyne приложение с уникальным ID
	a := fyneApp.NewWithID("report_app")

	// Загружаем конфигурацию (тема)
	cfg := config.Load()

	// Устанавливаем тему из конфига
	if cfg.Theme == "light" {
		a.Settings().SetTheme(theme.LightTheme())
	} else {
		a.Settings().SetTheme(theme.DarkTheme())
	}

	// Создаём окно
	w := a.NewWindow("Report Tracker v1.0")

	// Попытка загрузить иконку
	iconData, err := os.ReadFile("icon.png")
	if err == nil {
		w.SetIcon(fyne.NewStaticResource("icon.png", iconData))
	}

	// Создаём хранилище и сервис
	store := storage.NewFileStorage()
	service := service.NewStatsService(store)

	// Создаём UI
	content := ui.NewMainWindow(w, service, a, &cfg)
	w.SetContent(content)

	// Размер окна и центрирование
	w.Resize(ui.DefaultWindowSize())
	w.CenterOnScreen()

	// Запуск приложения
	w.ShowAndRun()
}

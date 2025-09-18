package core

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	*fiber.App
	db  *gorm.DB
	cfg *CoreConfig
	err error
}

func New(fiberApp *fiber.App, cfg ...CoreConfig) *App {
	app := &App{
		App: fiberApp,
	}

	if len(cfg) != 0 {
		c := cfg[0]
		app.cfg = &c
	} else {
		app.LoadConfig()
	}

	return app
}

func (app *App) hasError() bool {
	return app.err != nil
}

func (app *App) Run(host, port string) error {
	app.InitDB()
	app.LoadConfig()
	app.RegisterRoutes()

	if app.hasError() {
		return app.err
	}

	return app.Listen(host + ":" + port)
}

package console

import (
	"MyAPI/config"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

// New create new iris application for API builder
func New() *iris.Application {
	app := iris.New()
	return app
}

// Party is used for create new routes
func Party(app *iris.Application, dir string) router.Party {
	return app.Party(dir)
}

// Listen is used for running application
func Listen(app *iris.Application) {
	port := config.GetEnv("SERVER_PORT", "8080")
	app.Listen(fmt.Sprintf(":%s", port))
}

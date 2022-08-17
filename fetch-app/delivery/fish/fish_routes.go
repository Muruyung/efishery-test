package fish

import (
	"MyAPI/console"

	"github.com/kataras/iris/v12"
)

func Init(app *iris.Application) {
	fishAPI := console.Party(app, "/fish")
	{
		fishAPI.Use(iris.Compression)
		fishAPI.Get("/", ReadList)
		fishAPI.Get("/aggregate", Aggregate)
	}
}

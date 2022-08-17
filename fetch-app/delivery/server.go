package delivery

import (
	"MyAPI/adapters"
	"MyAPI/console"
	"MyAPI/controller"
	"MyAPI/delivery/fish"
	"fmt"
)

func Run() {
	app := console.New()
	fish.Init(app)

	adapters, err := adapters.Init()
	if err != nil {
		fmt.Println(err)
	}

	controller.InitController(adapters)
	console.Listen(app)
}

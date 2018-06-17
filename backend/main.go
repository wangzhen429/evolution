package main

import (
	"quant/backend/bootstrap"
	"quant/backend/bootstrap/database"
	"quant/backend/bootstrap/route"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("quant", "wuzhongyang@wzy.com")
	app.Bootstrap()
	app.Configure(database.Configure, route.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
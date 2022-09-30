package main

import (
	"main/findRestro"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func setupRoute(app *fiber.App) {
	app.Post("/findRestro", findRestro.FindR)
	app.Get("/", findRestro.FindForm)
}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoute(app)

	app.Listen(":3000")
}

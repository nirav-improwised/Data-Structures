package main

import (
	"main/findRestro"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func setupRoute(app *fiber.App) {
	app.Post("/restro", findRestro.AddR)
	app.Get("/restro", findRestro.FindR)
	app.Delete("/restro", findRestro.DeleteR)
	app.Put("/restro", findRestro.UpdateR)
}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoute(app)

	app.Listen(":3000")
}

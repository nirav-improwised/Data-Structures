package main

import (
	"main/findRestro"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	setupRoute(app)

	app.Listen(":5500")
}

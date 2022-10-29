package main

import (
	"hng/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RoutesAPI(app)

	app.Listen(":8080")

}

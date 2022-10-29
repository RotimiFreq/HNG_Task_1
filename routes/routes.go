package routes

import (
	"hng/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoutesAPI(app *fiber.App) {
	app.Get("/intern_details", controllers.UserAPI)

}

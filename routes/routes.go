package routes

import (
	"hng/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoutesAPI(app *fiber.App) {
	app.Get("/", controllers.UserAPI)
	app.Post("/arith", controllers.Arith_Operator)

}

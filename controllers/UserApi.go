package controllers

import (
	"context"
	"hng/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UserAPI(c *fiber.Ctx) error {

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	intern := models.InternDetails{
		SlackUsername: "Rotimihz",
		Backend:       true,
		Age:           23,
		Bio:           "Calm and not lazy",
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{"data": intern})

}

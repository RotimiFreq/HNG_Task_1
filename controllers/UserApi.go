package controllers

import (
	"context"
	"hng/models"

	//"hng/util"
	"net/http"
	//"strconv"
	//"strings"
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
func Arith_Operator(c *fiber.Ctx) error {

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	var req_body models.Req_body
	var Response_Body models.Response_Body

	if parse_err := c.BodyParser(&req_body); parse_err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"data": parse_err.Error()})

	}
	var Operation = []string{"addition", "substraction", "multiplication"}
	op := req_body.Operation_type
	x := req_body.X
	y := req_body.Y
	Response_Body.Username = "rotimihz"

	switch op {
	case Operation[0]:
		Response_Body.Operation_type = Operation[0]
		Response_Body.Result = x + y

	case Operation[1]:
		Response_Body.Operation_type = Operation[1]
		Response_Body.Result = x - y
	case Operation[2]:
		Response_Body.Operation_type = Operation[2]
		Response_Body.Result = x * y

	}

	return c.Status(http.StatusOK).JSON(Response_Body)

}

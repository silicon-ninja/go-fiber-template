package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorldController(c *fiber.Ctx) error {

	return c.SendString("Hello, World! Mamamamamam\n")

}

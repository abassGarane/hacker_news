package controllers

import "github.com/gofiber/fiber/v2"

type HomeController struct{}

func (h HomeController) Index(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{
		"Title": "Home page",
	}, "layouts/index")
}

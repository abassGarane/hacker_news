package controllers

import "github.com/gofiber/fiber"

type BlogController struct {
}

func (b *BlogController) Create(c *fiber.Ctx) error {
	return c.Render("pages/blogs/new", fiber.Map{
		"message": "successfully created",
	})
}

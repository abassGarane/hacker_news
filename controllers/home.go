package controllers

import (
	"hacker_news/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HomeController struct {
	Repo *repository.Repository
}

func (h HomeController) Index(c *fiber.Ctx) error {
	items, err := h.Repo.GetAllNews()
	log.Debug(items[0])
	if err != nil {
		c.Redirect("/404", http.StatusInternalServerError)
	}
	return c.Render("pages/index", fiber.Map{
		"items": items,
	}, "layouts/index")
}

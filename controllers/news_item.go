package controllers

import (
	"hacker_news/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type NewsItemController struct {
	DB *repository.Repository
}

func (n *NewsItemController) Index(c *fiber.Ctx) error {
	return c.Render("pages/news_item", fiber.Map{
		"Title": "News Item",
	}, "layouts/index")
}

func (n *NewsItemController) Submit(c *fiber.Ctx) error {
	return c.Render("/pages/submit", fiber.Map{}, "layouts/index")
}

func (n *NewsItemController) View(c *fiber.Ctx) error {
	idStr := c.Params("id", "1")
	log.Info("View item reached")
	id, err := strconv.ParseInt(idStr, 16, 64)
	if err != nil {
		c.Redirect("/")
	}
	log.Infof("view item id %d", id)
	item, err := n.DB.GetOneById(id)
	log.Debug(err)
	if err != nil {
		c.Redirect("/")
	}
	log.Debug(item)
	return c.Render("pages/news_item", fiber.Map{
		"item": item,
	}, "layouts/index")
}

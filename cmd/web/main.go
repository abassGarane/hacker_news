package main

import (
	"fmt"
	"hacker_news/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("hello world!")
}

func initServer() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Static("/static", "./static")

	homeContr := controllers.HomeController{}

	app.Get("/", homeContr.Index)

	log.Fatal(app.Listen(":8080"))
}

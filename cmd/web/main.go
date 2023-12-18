package main

import (
	"hacker_news/controllers"
	"hacker_news/repository"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	initServer()
}

func initServer() {
	slog.Info("Creating db")
	dsn := os.Getenv("DSN")
	slog.Debug(dsn)
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/hacker_news?sslmode=disable"
		slog.Error("could not find dsn")
	}
	db, err := repository.New(dsn)
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Creating a server with fiber")
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Middlewares
	app.Use(recover.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Serve static files
	app.Static("/", "./static")

	homeContr := controllers.HomeController{
		Repo: db,
	}

	app.Get("/", homeContr.Index)

	newsItemContr := controllers.NewsItemController{DB: db}

	app.Get("/news/:id", newsItemContr.View)
	app.Get("/news", newsItemContr.Index)

	slog.Info("Running server on port 8080")

	log.Fatal(app.Listen(":8080"))
}

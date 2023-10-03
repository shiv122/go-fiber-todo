package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shiv122/go-todo/config"
	"github.com/shiv122/go-todo/connection"
	"github.com/shiv122/go-todo/routes"
)

func main() {
	connection.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// prometheus := fiberprometheus.New("my-service-name")
	// prometheus.RegisterAt(app, "/metrics")
	// app.Use(prometheus.Middleware)

	routes.SetupApiRoute(app)

	// migration.Migrate()

	app.Listen(":" + config.App.Port)
}

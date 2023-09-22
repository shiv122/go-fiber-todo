package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shiv122/go-todo/config"
	"github.com/shiv122/go-todo/connection"
	"github.com/shiv122/go-todo/migration"
	"github.com/shiv122/go-todo/routes"
)

func main() {
	app := fiber.New()

	// prometheus := fiberprometheus.New("my-service-name")
	// prometheus.RegisterAt(app, "/metrics")
	// app.Use(prometheus.Middleware)

	connection.ConnectDB()

	migration.Migrate()

	routes.SetupApiRoute(app)

	// log.Info(config.App)

	app.Listen(":" + config.App.Port)
}

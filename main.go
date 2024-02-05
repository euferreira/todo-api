package main

import (
	"github.com/euferreira/pkg/configurations"
	"github.com/euferreira/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := configurations.CreateClient()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	routes.AddRoutes(app)

	app.Listen(":3000")
}

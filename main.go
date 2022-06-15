package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_hello_world_routes "github.com/silicon-ninja/m/src/routes"
	_project_utils "github.com/silicon-ninja/m/src/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "@silicon-ninja", // add cusViews: engine,
	})

	server_environment := _project_utils.Config("environment")
	server_port := _project_utils.Config("port")

	if server_environment == "development" {
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
		_project_utils.DebugLogger("Server is running in development mode")
	}

	_hello_world_routes.HelloWorldRoutes(app)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("█▀▀ █▀█   █▀▀ █ █▄▄ █▀▀ █▀█   █▀ ▀█▀ ▄▀█ █▀█ ▀█▀ █▀▀ █▀█")
	fmt.Println("█▄█ █▄█   █▀░ █ █▄█ ██▄ █▀▄   ▄█ ░█░ █▀█ █▀▄ ░█░ ██▄ █▀▄")
	fmt.Println("---------------------------------------------")
	fmt.Println("")
	_project_utils.InfoLogger("Server is running in " + server_environment + " mode on port " + server_port)
	err := app.Listen(":" + server_port)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}

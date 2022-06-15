package hello_world_routes

import (
	"github.com/gofiber/fiber/v2"
	_hello_world_controller "github.com/silicon-ninja/m/src/controllers"
)

func HelloWorldRoutes(hellorouter *fiber.App) {
	hellorouter.Get("/helloworld", _hello_world_controller.HelloWorldController)
}

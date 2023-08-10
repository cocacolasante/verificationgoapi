package server

import (
	"verificationapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer() *fiber.App {
	app := fiber.New()

	app.Get("/userspostsall/:address", handlers.GetUsersPostsAll)

	app.Get("/userspost/:address/:postnum", handlers.GetSingleUsersPost)

	app.Get("/mintfee", handlers.GetMintFee)

	return app
}



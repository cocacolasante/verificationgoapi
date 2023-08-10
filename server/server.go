package server

import (
	"verificationapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer() *fiber.App {
	app := fiber.New()

	// get routes for verification contract
	app.Get("/userspostsall/:address", handlers.GetUsersPostsAll)

	app.Get("/userspost/:address/:postnum", handlers.GetSingleUsersPost)

	app.Get("/mintfee", handlers.GetMintFee)

	app.Get("/getprofile/:address", handlers.GetProfile)

	return app
}



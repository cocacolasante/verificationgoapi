package main

import (
	"fmt"
	"verificationapi/client"
	"verificationapi/contractinterfaces"
	"verificationapi/handlers"

	"github.com/gofiber/fiber/v2"
	// "github.com/ethereum/go-ethereum/common"
)

func main(){
	newClient := client.ConnectToClient()

	fmt.Println("Connected to", newClient)

	fee := contractinterfaces.GetFee()

	fmt.Println(fee)

	app := fiber.New()
	

	// addy := common.HexToAddress("0x11273F391609BF4C05CA23c6aD29D919a71dc37E")

	// postsSlice := contractinterfaces.GetAllUsersPosts(addy)

	// fmt.Println(postsSlice)

	// postTwo := contractinterfaces.GetSpecificPost(addy, 2)

	// fmt.Println(postTwo)

	app.Get("/userspostsall/:address", handlers.GetUsersPostsAll)

	app.Listen(":3000")
	fmt.Println("App listening on port 3000")
	
}
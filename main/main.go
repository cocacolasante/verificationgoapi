package main

import (
	"fmt"
	"verificationapi/client"
	"verificationapi/contractinterfaces"
	// "verificationapi/handlers"
	"verificationapi/server"

	// "github.com/ethereum/go-ethereum/common"
)

func main(){
	newClient := client.ConnectToClient()

	fmt.Println("Connected to", newClient)

	fee := contractinterfaces.GetFee()

	fmt.Println(fee)

	
	
	apiServer := server.StartServer()
	
	apiServer.Listen(":3000")
	fmt.Println("App listening on port 3000")

	
	

	
	
}
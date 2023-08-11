package main

import (
	"fmt"
	// "verificationapi/handlers"
	"verificationapi/server"

	// "github.com/ethereum/go-ethereum/common"
)

func main(){
	apiServer := server.StartServer()
	
	apiServer.Listen(":3000")
	fmt.Println("App listening on port 3000")

	
	

	
	
}
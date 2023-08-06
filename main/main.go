package main

import (
	"fmt"
	"verificationapi/client"
)

func main(){
	newClient := client.ConnectToClient()

	fmt.Println("Connected to", newClient)
}
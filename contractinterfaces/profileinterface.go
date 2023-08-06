package contractinterfaces

import (
	"log"
	"os"
	"verificationapi/client"
	profile "verificationapi/profilecontract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)



func GetOwner() common.Address {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	client := client.ConnectToClient()

	profileAddy := os.Getenv("PROFILE_ADDRESS")

	profileAddress := common.HexToAddress(profileAddy)

	profileContract, err := profile.NewProfile(profileAddress, &client)

	if err != nil{
		log.Fatal("Error fatching profile contract", err)
	}


	deployerAddy := os.Getenv("DEPLOYER_PUBLIC")
	deployerAddress := common.HexToAddress(deployerAddy)
	owner, err := profileContract.Owner(&bind.CallOpts{
		From: deployerAddress,
	})
	if err != nil {
		log.Fatal("Error returning owner", err)
	}

	return owner
}


func GetProfileByAddress(address common.Address) common.Address{
	
}
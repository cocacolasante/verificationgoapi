package contractinterfaces

import (
	"log"
	"os"
	"verificationapi/client"
	profile "verificationapi/profilecontract"
	"verificationapi/web3types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func getProfileContract(c *ethclient.Client) *profile.Profile {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	
	profileAddy := os.Getenv("PROFILE_ADDRESS")
	profileAddress := common.HexToAddress(profileAddy)
	profileContract, err := profile.NewProfile(profileAddress, c)

	if err != nil{
		log.Fatal("Error fatching profile contract", err)
	}

	return profileContract

}

func GetOwner() common.Address {
	

	client := client.ConnectToClient()

	profileContract := getProfileContract(&client)
	

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


func GetProfileByAddress(address common.Address) web3types.ProfileStruct{
	client := client.ConnectToClient()

	profileContract := getProfileContract(&client)

	profile, err := profileContract.Profiles(&bind.CallOpts{}, address)

	if err != nil{
		log.Fatal("Error getting profile", err)

	}
	return profile
}


// post request // will probably not use it for post requests -- unsecure

// func CreateNewProfile(username string, qrCode string, sender common.Address) {
// 	client := client.ConnectToClient()

// 	newProfile := web3types.ProfileStruct{ProfileAddress: sender, Username: username, ProfileQRCode: qrCode}

	

// }
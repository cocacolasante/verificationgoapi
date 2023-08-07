package contractinterfaces

import (
	"log"
	"math/big"
	"os"
	"verificationapi/client"
	verification "verificationapi/verificationcontract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// getFee getAllUsersPost getSpecificPost getAdmin


func getVerificationContract(c *ethclient.Client) *verification.Verification{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error fetching smart contract", err)

	}

	verificationAddy := os.Getenv("VERIFICATION_ADDRESS")
	verificationAddress := common.HexToAddress(verificationAddy)

	verificationContract, err := verification.NewVerification(verificationAddress, c)
	if err != nil {
		log.Fatal("Error getting Verification contract", err)
	}
	return verificationContract
}

func GetFee() *big.Int {
	c := client.ConnectToClient()
	
	verificationContract := getVerificationContract(&c)
	deployerAddy := os.Getenv("DEPLOYER_PUBLIC")
	deployerAddress := common.HexToAddress(deployerAddy)

	fee, err := verificationContract.GetFee(&bind.CallOpts{
		From: deployerAddress,
	})
	if err != nil {
		log.Fatal("Cannot get fee", err)

	}
	return fee
}

func GetAdmin() common.Address{
	c := client.ConnectToClient()
	verificationContract := getVerificationContract(&c)

	admin, err := verificationContract.GetAdmin(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Error fetching admin")
	}

	return admin
}

func GetAllUsersPosts(address common.Address)[]verification.VerificationPost{
	c := client.ConnectToClient()
	
	verificationContract := getVerificationContract(&c)


	usersPosts, err := verificationContract.GetAllUsersPost(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal("Error fetching posts", err)

	}

	return usersPosts
} 

func GetSpecificPost(address common.Address, postNum uint) verification.VerificationPost {
	c := client.ConnectToClient()
	
	verificationContract := getVerificationContract(&c)

	var biPost big.Int
	biPost.SetUint64(uint64(postNum))


	specificPost, err := verificationContract.GetSpecificPost(&bind.CallOpts{}, address, &biPost)
	if err != nil {
		log.Fatal("Error fetching post", err)
	}

	return specificPost
}


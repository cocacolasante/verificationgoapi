package handlers

import (
	"fmt"
	"strconv"
	"verificationapi/contractinterfaces"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
)

// handlers for verification contract
func GetUsersPostsAll(c *fiber.Ctx) error {
	
	addy := c.Params("address")

	address := common.HexToAddress(addy)
	
	posts, err := contractinterfaces.GetAllUsersPosts(address)
	if err != nil {
		return err
	}


	c.JSON(posts)
	

	return nil


}

func GetSingleUsersPost(c *fiber.Ctx) error {
	addy := c.Params("address")
	postNum := c.Params("postnum")
	

	address := common.HexToAddress(addy)

	postNumber, err := strconv.ParseUint(postNum, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return err
	}
	

	post := contractinterfaces.GetSpecificPost(address, postNumber)

	c.Status(200).JSON(post)

	return nil

}

func GetMintFee(c *fiber.Ctx) error {
	mintFee, err := contractinterfaces.GetFee()
	if err != nil {
		c.Status(500).JSON(err)
		return err
	}

	c.Status(200).JSON(mintFee)

	return nil
}


// handlers for profile contract

func GetProfile(c *fiber.Ctx) error {
	addy := c.Params("address")
	address := common.HexToAddress(addy)
	profile := contractinterfaces.GetProfileByAddress(address)

	
	c.Status(200).JSON(profile)

	return nil
}


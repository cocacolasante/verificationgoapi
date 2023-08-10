package handlers

import (
	"verificationapi/contractinterfaces"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
)

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
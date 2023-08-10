package handlers

import (
	"fmt"
	"strconv"
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

func GetSingleUsersPost(c *fiber.Ctx) error {
	addy := c.Params("address")
	postNum := c.Params("postnum")
	

	address := common.HexToAddress(addy)

	postNumber, err := strconv.ParseUint(postNum, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}
	

	post := contractinterfaces.GetSpecificPost(address, postNumber)

	c.JSON(post)

	return nil

}

func GetMintFee(c *fiber.Ctx) error {
	mintFee := contractinterfaces.GetFee()
	
	c.JSON(mintFee)

	return nil
}
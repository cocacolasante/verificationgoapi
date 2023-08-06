package client

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func ConnectToClient() ethclient.Client{
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mumbaiUrl := os.Getenv("MUMBAI_URL")

	client, err := ethclient.DialContext(context.Background(), mumbaiUrl)
	if err != nil {
		log.Fatal("Cannot connect to client")
	}
	defer client.Close()

	return *client
}

func GetChainId(c ethclient.Client) *big.Int{
	chainId, err := c.ChainID(context.Background())
	if err != nil {
		log.Fatal("Cannot get chain id")
	}
	return chainId
}

func GetBlock(c ethclient.Client) *types.Block{
	block, err := c.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal("Cannot get block")
	}
	return block
}

func GetBalance(c ethclient.Client, addr string) *big.Int{
	address := common.HexToAddress(addr)

	block := GetBlock(c)

	balance, err := c.BalanceAt(context.Background(), address, block.Number())
	if err != nil{
		log.Fatal("cannot get balance")
	}
	return balance
}


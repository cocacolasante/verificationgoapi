package web3types

import 	"github.com/ethereum/go-ethereum/common"

type ProfileStruct struct {
	ProfileAddress common.Address
	Username       string
	ProfileQRCode  string
}
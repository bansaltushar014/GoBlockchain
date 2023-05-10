package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bansaltushar014/golangBlockchain/wallet"
)

func callWallet(nodeId string) {
	p := wallet.CreateWallet(nodeId)
	p.AddWallet()
	p.SaveFile(nodeId)
}

func getWallet(address string, nodeId string) {
	fmt.Println("Get Wallet!")
	p := wallet.CreateWallet(nodeId)
	p.LoadFile(nodeId)
	fmt.Println(p.Wallets[address].PublicKey)
	fmt.Println(p)
	wallet.ValidateAddress(address)
}

func getAllAddress(nodeId string) {
	fmt.Println("Get All Addresses!")
	p := wallet.CreateWallet(nodeId)
	addresses := p.GetAllAddresses(nodeId)
	fmt.Println(addresses)
}

func main() {

	switch os.Args[1] {
	case "createWallet":
		callWallet(os.Args[2])
	case "getWallet":
		getWallet(os.Args[2], os.Args[3])
	case "getAllAddress":
		getAllAddress(os.Args[2])
	default:
		// cli.printUsage()
		runtime.Goexit()
	}
}

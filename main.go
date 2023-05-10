package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bansaltushar014/golangBlockchain/wallet"
)

func callWallet() {
	p := wallet.CreateWallet()
	p.AddWallet()
	p.SaveFile("3000")
}

func getWallet(address string) {
	fmt.Println("Get Wallet!")
	p := wallet.CreateWallet()
	p.LoadFile("3000")
	fmt.Println(p.Wallets[address].PublicKey)
	fmt.Println(p)
	wallet.ValidateAddress(address)
}

func main() {

	switch os.Args[1] {
	case "createWallet":
		callWallet()
	case "getWallet":
		getWallet(os.Args[2])
	default:
		// cli.printUsage()
		runtime.Goexit()
	}
}

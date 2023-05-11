package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bansaltushar014/golangBlockchain/blocks"
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

func genesisBlock() {
	fmt.Println("Create the Genesis Block!")
	blocks.GenesisBlock()
}

func AddBlock() {
	fmt.Println("Continue Blockchain!")
	b := blocks.ContinueBlockChain()
	b.AddBlock([]byte("Second Block"))

	fmt.Println(b)
}

func iter() {
	fmt.Println("Iterate through Blockchain!")
	b := blocks.ContinueBlockChain()
	iter := &blocks.BlockChainIterator{b.LastHash, b.Database}
	s := iter.Iter()
	iter2 := &blocks.BlockChainIterator{s, b.Database}
	s2 := iter2.Iter()
	fmt.Println(s2)
}

func main() {

	switch os.Args[1] {
	case "createWallet":
		callWallet(os.Args[2])
	case "getWallet":
		getWallet(os.Args[2], os.Args[3])
	case "getAllAddress":
		getAllAddress(os.Args[2])
	case "genesisBlock":
		genesisBlock()
	case "AddBlock":
		AddBlock()
	case "iter":
		iter()
	default:
		// cli.printUsage()
		runtime.Goexit()
	}
}

package blocks

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bansaltushar014/golangBlockchain/wallet"
	"github.com/dgraph-io/badger"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

const dbFile = "./tmp/blocks/MANIFEST"
const dbPath = "./tmp/blocks"

func DBexists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func ContinueBlockChain() *BlockChain {
	if DBexists() == false {
		fmt.Println("No existing blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	db, err := badger.Open(badger.DefaultOptions(dbPath))
	wallet.Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		wallet.Handle(err)

		err = item.Value(func(val []byte) error {
			lastHash = append([]byte{}, val...)
			return nil
		})

		return err
	})
	wallet.Handle(err)

	chain := BlockChain{lastHash, db}

	return &chain
}

func (chain *BlockChain) AddBlock(data []byte) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		wallet.Handle(err)
		// lastHash, err = item.Value()
		err = item.Value(func(val []byte) error {
			lastHash = append([]byte{}, val...)
			return nil
		})

		return err
	})
	wallet.Handle(err)
	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		wallet.Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})

	wallet.Handle(err)
}

func GenesisBlock() {
	newBlock := CreateBlock([]byte("First Block"), []byte("nil"))
	db, err := badger.Open(badger.DefaultOptions(dbPath))

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		wallet.Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		// lastHash := newBlock.Hash

		return err
	})
	wallet.Handle(err)
}

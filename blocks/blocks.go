package blocks

import (
	"bytes"
	"encoding/gob"

	"github.com/bansaltushar014/golangBlockchain/wallet"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data []byte, PrevHash []byte) *Block {
	b := Block{nil, data, PrevHash, 0}
	pow := NewProof(&b)
	nonce, hash := pow.Run()

	b.Hash = hash[:]
	b.Nonce = nonce

	return &b
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	wallet.Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	wallet.Handle(err)

	return &block
}

package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/HoYaStudy/Go_Study/hcoin/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

var ErrNotFound = errors.New("Block not found")

func FindBlock(hash string) (*Block, error) {
	blockBytes := dbStorage.FindBlock(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func createBlock(prevHash string, height int, difficulty int) *Block {
	block := &Block{
		Hash: "", PrevHash: prevHash, Height: height, Difficulty: difficulty, Nonce: 0,
	}
	block.Transactions = Mempool().TxToConfirm()
	block.mine()
	persistBlock(block)
	return block

}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			return
		} else {
			b.Nonce++
		}
	}
}

func persistBlock(b *Block) {
	dbStorage.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

var Blockchain []Block
type Block struct {
	Index int
	TimeStamp string
	BPM int
	Hash string
	PreHash string
	Validator string
}

// SHA256 hasing
// calculateHash is a simple SHA256 hashing function
func CalculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CalculateBlockHash(block Block) string {
	record := string(block.Index) + block.TimeStamp + string(block.BPM) + block.PreHash
	return CalculateHash(record)
}

func GenerateBlock(oldBlock Block, BPM int) (Block, error) {

	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.String()
	newBlock.BPM = BPM
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = CalculateBlockHash(newBlock)

	return newBlock, nil
}

func GenerateBlockForPos(oldBlock Block, BPM int,addr string) (Block, error) {

	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.String()
	newBlock.BPM = BPM
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = CalculateBlockHash(newBlock)
	newBlock.Validator=addr;
	return newBlock, nil
}


func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PreHash {
		return false
	}

	if CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// make sure the chain we're checking is longer than the current blockchain
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
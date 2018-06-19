package main

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
}

func calculateHash(block Block) string{
	record:=string(block.Index)+block.TimeStamp+string(block.BPM)+block.PreHash
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block,BPM int)(Block,error){
	 var newBlock Block
	 t:=time.Now()
	 newBlock.Index=oldBlock.Index+1
	 newBlock.TimeStamp=t.String()
	 newBlock.BPM=BPM
	 newBlock.PreHash=oldBlock.Hash
	 newBlock.Hash=calculateHash(newBlock)
	 return newBlock,nil
}

func isBlockValid(newBlock,oldBlock Block) bool{
	if oldBlock.Index+1 !=newBlock.Index{
		return  false
	}

	if oldBlock.Hash!=newBlock.PreHash {
		return false
	}

	if calculateHash(newBlock)!= newBlock.Hash{
		return false
	}

	return true
}

func replaceChain(newblocks []Block){
	if len(newblocks)>len(Blockchain){
		Blockchain=newblocks
	}
}

func main() {
	
}

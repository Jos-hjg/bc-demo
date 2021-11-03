package blockchain

type BlockChain struct {
	lastBlockHash Hash
	blocks        map[Hash]*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: map[Hash]*Block{},
	}
}

func (bc *BlockChain) AddGensisBlock() *BlockChain {
	if bc.lastBlockHash != "" {
		return bc
	}
	return bc.AddBlock("The Gensis Block.")
}

func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	b := NewBlock(bc.lastBlockHash, txs)
	bc.blocks[b.hashCurr] = b
	bc.lastBlockHash = b.hashCurr
	return bc
}

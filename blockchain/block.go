package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash = string

const nodeVersion = 0

//主体
type Block struct {
	header    BlockHeader
	txs       string //交易列表
	txCounter int    //交易计数器
	hashCurr  Hash   //当前区块Hash值缓存
}

//头
type BlockHeader struct {
	version        int       //版本号
	hashPrevBlock  Hash      //前一个区块的hash
	hashMerkleRoot Hash      //默克尔树的hash节点
	time           time.Time //区块创建时间
	bits           int       //难度相关
	nonce          int       //挖矿相关
}

func (bh *BlockHeader) Stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(),
		bh.bits,
		bh.nonce,
	)
}

func (b *Block) SetHashCurr() *Block {
	headerStr := b.header.Stringify()
	b.hashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte(headerStr)))
	return b
}

func NewBlock(prevHash Hash, txs string) *Block {
	b := &Block{
		header: BlockHeader{
			version:       nodeVersion,
			hashPrevBlock: prevHash,
			time:          time.Now(),
		},
		txs:       txs,
		txCounter: 1,
	}
	b.SetHashCurr()
	return b
}

package WDC

/**
  区块链相关操作
 */
type Blockchain struct {
	blocks[] *Block
}

//新增区块
func (blockchain *Blockchain) AddBlocktoChain(data string){

     preBlock:=blockchain.blocks[len(blockchain.blocks)-1] //获取上一个区块信息

     newBlocks:=NewBlock(data,preBlock.Hash)

	// 2. 将区块添加到Blocks里面
	 blockchain.blocks=append(blockchain.blocks,newBlocks)


}

// 创建一个带有创世区块节点的区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

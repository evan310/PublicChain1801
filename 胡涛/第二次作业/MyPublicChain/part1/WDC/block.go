package WDC

/**
区块
 */
import (
	"time"
	"strconv"
	"bytes"
	"fmt"
	"crypto/sha256"
)

//区块相关属性
type Block struct {
	// 时间戳，创建区块时的时间
	Timestamp     int64
	// 上一个区块的hash，父hash
	PrevBlockHash []byte
	// Data，交易数据
	Data []byte
	// Hash 当前区块的Hash
	Hash []byte
	//Notice 挖矿难度 随机数
	Nonce int

}

// 工厂方法 实例化一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	// 创建区块
	block:=&Block{time.Now().Unix(),prevBlockHash,[]byte(data),[]byte{},0}
    //设置hash值
   // block.SetHash()


   //工作量证明、生成区块，并添加到区块链
     pow:=CreateProofOfWork(block)

	// Run()执行一次工作量证明
	nonce, hash := pow.Run()

	block.Nonce=nonce

	block.Hash=hash[:]

	isValid := pow.validate()

	fmt.Println(isValid)


    return block
}

func (block *Block) SetHash()  {

	// 1. 将时间戳转字节数组
	//（1）将int64转字符串
	// 第二个参数的范围为2 ~ 36，代表进制
	timeString := strconv.FormatInt(block.Timestamp,2)
	//fmt.Println(timeString)
	//（2）将字符串转字节数组
	timestamp := []byte(timeString)
	// 2. 将除了Hash以外的其他属性，以字节数组的形式全拼接起来

	//fmt.Println([][]byte{block.PrevBlockHash, block.Data, timestamp})
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp},[]byte{})
	//fmt.Println("fmt.Println(headers)")
	fmt.Println(headers)
	// 3. 将拼接起来数据进行256 hash
	hash := sha256.Sum256(headers)
	// 4. 将hash赋给Hash属性字节
	block.Hash = hash[:]
}



// 创建创世区块，并返回创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genenis Block",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
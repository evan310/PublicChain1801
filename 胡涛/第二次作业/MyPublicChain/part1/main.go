package main
import (
	"./WDC"
)

func main()  {

	blockchain := WDC.NewBlockchain()

	blockchain.AddBlocktoChain("add block1 to blockchain")

	blockchain.AddBlocktoChain("add block2 to blockchain")

	blockchain.AddBlocktoChain("add block3 to blockchain")


	/*for _, blockd := range blockchain.blocks{

		fmt.Printf("Data：%s \n",string(blockd.Data))
		fmt.Printf("PrevBlockHash：%x \n",blockd.PrevBlockHash)
		fmt.Printf("Timestamp：%s \n",time.Unix(blockd.Timestamp, 0).Format("2018-06-29 03:04:05 PM") )
		fmt.Printf("Hash：%x \n",blockd.Hash)
		fmt.Printf("Nonce：%d \n",blockd.Nonce)

		fmt.Println();
	}*/
}
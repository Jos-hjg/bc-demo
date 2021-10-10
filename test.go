package main

import (
	"fmt"
	"github.com/Jos-hjg/bc-demo/blockchain"
)

func main() {
	b := blockchain.NewBlock("", "Gensis Block.")
	fmt.Println(b)
}

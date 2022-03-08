package main

import (
	"fmt"
	"log"

	"opencrud.com/blockchain/utils"
)

func init() {
	log.SetPrefix("Cmd Server:")
}

func main() {
	// fmt.Println(utils.FindNeighbors("127.0.0.1", 5000, 0, 3, 5000, 5003))
	fmt.Println(utils.GetHost())
}

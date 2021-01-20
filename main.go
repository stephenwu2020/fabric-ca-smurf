package main

import (
	"fmt"

	"github.com/hyperledger/fabric-ca/lib"
)

func main() {
	fmt.Println("Smurf runing...")
	c := lib.ClientConfig{}
	fmt.Println(c)
}

package main

import (
	"fmt"

	"example.com/go/crypto/api"
)


func main() {
	rate, err := api.GetRate("BTC")
	fmt.Println(rate, err)
}

// First think about we wart to structure the data. 
// create a package (folder)
// See datatypes
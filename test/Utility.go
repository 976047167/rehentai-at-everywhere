package main

import (
	"base"
	"fmt"
)

// this is the unittest of `Utility`
func main() {
	fmt.Println("GetServerTime", base.GetServerTime())
	fmt.Println("BinaryToHex", base.BinaryToHex([]byte("binary2hex")))
	fmt.Println("GetSHA1String", base.GetSHA1String("string2hash"))
	client := base.NewClient()
	client.Run()
}

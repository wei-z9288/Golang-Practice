package main

import (
	"fmt"
	algo "fops/algo"
)

func main() {
	var a algo.Algorithm = algo.GetMD5()
	data := []byte("hello world")
	fmt.Printf("Type: %s\n", a.GetType())
	fmt.Printf("Checksum: %s\n", a.GetChecksum(data))

	var b algo.Algorithm = algo.GetSHA1()
	fmt.Printf("Type: %s\n", b.GetType())
	fmt.Printf("Checksum: %s\n", b.GetChecksum(data))
}

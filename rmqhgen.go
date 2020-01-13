package main

import (
	"flag"
	"fmt"
)

func main() {
	var password, algo string

	flag.StringVar(&password, "password", "", "The password to hash")
	flag.StringVar(&algo, "algorithm", "sha256", "The hash algorithm to use (either sha256 or sha512)")

	flag.Parse()

	fmt.Println("Password to hash:", password)
	fmt.Println("Hash algorithm:", algo)
}

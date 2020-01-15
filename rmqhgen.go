package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	var args = flag.Args()

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Expected exactly one argument (args: %v)\n", args)
		os.Exit(1)
	}

	fmt.Println(generate(args[0]))
}

func generate(password string) string {
	var salt = generateSalt()
	var saltedPass = append(salt, password...)
	var hash = sha256.Sum256(saltedPass)
	var saltedHash = append(salt, hash[:]...)

	return base64.StdEncoding.EncodeToString(saltedHash)
}

func generateSalt() []byte {
	var salt = make([]byte, 4)

	rand.Read(salt)

	return salt
}

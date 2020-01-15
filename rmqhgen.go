package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected either 'generate' or 'validate' subcommand")
		os.Exit(1)
	}

	var subcommandName = os.Args[1]
	var subcommandArgs = os.Args[2:]

	switch subcommandName {
	case "generate":
		if len(subcommandArgs) != 1 {
			fmt.Fprintf(os.Stderr, "Expected exactly one argument (args: %v)\n", subcommandArgs)
			os.Exit(1)
		}
		fmt.Println(generate(subcommandArgs[0]))
	case "validate":
		fmt.Println("TODO")
	default:
		fmt.Println("Expected either 'generate' or 'validate' subcommand")
		os.Exit(1)
	}
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

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("\033[0;31mYou should enter only 1 argument SHA256, SHA384, SHA512")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\033[0;32m If you want to skip entering the values then press CTRL + D instead\n")

	for scanner.Scan() {

		txt := scanner.Text()

		switch args[0] {
		case "SHA256":
			sha_256 := sha256.Sum256([]byte(txt))
			fmt.Printf("%x\n", sha_256)
		case "SHA512":
			sha_512 := sha512.Sum512([]byte(txt))
			fmt.Printf("%x\n", sha_512)
		case "SHA384":
			sha_384 := sha512.Sum384([]byte(txt))
			fmt.Printf("%x\n", sha_384)
		default:
			fmt.Println("Error... , try again: remember you should enter SHA256 or SHA384 or SHA512")
			os.Exit(1)
		}
	}
}

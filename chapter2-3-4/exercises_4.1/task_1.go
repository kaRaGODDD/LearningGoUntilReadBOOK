package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
	"os"
)

func main() {
    
    args := os.Args[1:]
    
    if len(args) != 2 {
        fmt.Fprintf(os.Stderr, "Amount of arguments should be 2\n")
        os.Exit(1)
    }

    c1 := sha256.Sum256([]byte(args[0]))
    c2 := sha256.Sum256([]byte(args[1]))
    
    cnt1, cnt2 := 0, 0

    for _, byte := range c1 {
        cnt1 += bits.OnesCount8(byte)
    }

    for _, byte := range c2 {
        cnt2 += bits.OnesCount8(byte)
    }

    fmt.Printf("Bits that not equal to zero in first hash: %d\nBits that not equal to zero in second hash: %d\n", cnt1, cnt2)
}

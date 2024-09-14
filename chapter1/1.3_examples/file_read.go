package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

    mp := make(map[string]int) 
       
    data, err := os.ReadFile("text.txt")

    if err != nil {
        fmt.Fprintf(os.Stderr, "The error was occured in program read file %v", err)
    }
    
    for _, str := range strings.Split(string(data), "\n") {
        mp[str]++
    }

    for cnt, str := range mp {
        fmt.Printf("Formate str %d : Formate cnt %s \n", str, cnt)


    }

}

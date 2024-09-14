package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {

    mp := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    for input.Scan() { 
        mp[input.Text()]++
    }
   
    for line, n := range mp {
        if n > 1 {
           fmt.Printf("%d\t%s\n", n, line) 
        }
    }

}

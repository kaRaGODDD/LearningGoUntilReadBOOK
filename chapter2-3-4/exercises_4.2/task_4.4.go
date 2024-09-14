package main

import (
    "fmt"
    "unicode"
)

func ajd(s []byte) []byte {
    idx := 0
    for i := 0; i < len(s); i++ {
        if i < len(s)-1 && unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
            continue 
        }
        s[idx] = s[i]
        idx++
    }
    if idx > 0 && unicode.IsSpace(rune(s[idx-1])) {
        s[idx-1] = ' '
    }
    return s[:idx]
}

func main() {
    s := []byte("  Hello   World  ")
    s = ajd(s)
    fmt.Println(string(s)) 
}

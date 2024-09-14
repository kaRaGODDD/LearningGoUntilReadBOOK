/*
Напишите функцию, которая без выделения дополнительной па
мяти удаляет все смежные дубликаты в срезе [ ] s t r i n g 
*/

package main

import (
    "fmt"
)

func delete_adj_el(s []string) []string {
    idx := 0
    for i := 0; i < len(s); i++ {
        if (s[i] != s[idx]) {
            idx++
            s[idx] = s[i]             
        }
    }
    return s[:idx + 1]
}

func main() {
    asdf := []string{"asdf", "asdf", "qwe"}
    asdf = delete_adj_el(asdf)
    fmt.Println(asdf)
}

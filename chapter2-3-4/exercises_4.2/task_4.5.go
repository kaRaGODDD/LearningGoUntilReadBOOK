/*
Перепишите функцию r e v e r s e так, чтобы она без выделения
дополнительной памяти обращала последовательность символов среза [ ] b y te , кото
рый представляет строку в кодировке UTF-8. Сможете ли вы обойтись без выделения
новой памяти?
*/

package main

import (
    "fmt"
    "unicode/utf8"
)

func reversae(a []byte) []byte {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    for i := 0; i < len(a); {
        _, size := utf8.DecodeRune(a[i:])
        for j := 0; j < size / 2; j++ {
            a[i+j], a[i+size-1-j] = a[i+size-1-j], a[i+j]
        }
        i += size
    }
    return a
}

func main() {
    a := []byte("Hello World")
    fmt.Println(a)
    reversae(a)
    fmt.Println(a)
}



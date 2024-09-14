/*
Перепишите функцию r e v e r s e так, чтобы вместо среза она ис
пользовала указатель на массив
*/

package main

import (
    "fmt"
)

func reverse (a *[]int) {
    for i, j := 0, len(*a) - 1; i < j; i, j = i + 1, j - 1 {
            (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
    }
}



func main() {
    src := []int{1,2,3,4,5}
    reverse(&src)
    fmt.Println(src)
}

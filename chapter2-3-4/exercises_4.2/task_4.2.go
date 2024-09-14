/*
Напишите версию функции r o t a t e , которая работает в один
проход
*/

package main

import (
    "fmt"
)

func rotate (a []int) {
    var t int = a[0]
    for i := 0; i < len(a) - 1; i++ {
        a[i] = a[i + 1] 
    } 
    a[len(a) - 1] = t
}

func main() {
    a := []int{1,2,3,4,5}
    rotate(a)
    fmt.Println(a)
    rotate(a)
    fmt.Println(a)
}

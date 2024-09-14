package main

import (
    "fmt"
)

type Currency int

const (
   USD Currency = iota 
   EUR
   GBP
   RUR
)


func main() {
    symbol := [...]string{USD: "$"}
    var a[3] int = [3]int{1,3,4}
    fmt.Println(symbol[USD], a[2])
}

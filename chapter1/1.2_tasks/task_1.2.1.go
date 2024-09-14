/*
    :Изменить программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды
*/

package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    
    args := strings.Split(os.Args[0], "/")
    
    fmt.Println(args[len(args) - 1])


}

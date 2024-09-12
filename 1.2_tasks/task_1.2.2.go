/*
    :Измените программу echo так, чтобы она выводила индекс и значение
    :каждого аргумента по одному аргументу в строке
*/

package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    
    arguments_slice := strings.Split(os.Args[0], "/")

    for ind, val := range arguments_slice {
        if ind != 0 {
            fmt.Println(ind, " : ", val)
        }
    }
}

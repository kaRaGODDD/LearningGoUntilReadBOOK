/*
Вызов функции i o . C o p y ( d s t , s r c ) выполняет чтение s r c и
запись в d s t . Воспользуйтесь ею вместо i o u t i l . R e a d A l l для копирования тела
ответа в поток o s . S t d o u t без необходимости выделения достаточно большого для
хранения всего ответа буфера. Не забудьте проверить, не произошла ли ошибка при
вызове i o . Сору.
*/

package main 

import (
    "fmt"
    "os"
    "io"
    "net/http"

)

func main() {

    resp, err := http.Get("https://youtube.com")
    
    if err != nil {
        fmt.Fprintf(os.Stderr,"Mistake %v", err)
    }

    wr, err := io.Copy(os.Stdout, resp.Body)
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "Msg %v", err)
    }

    fmt.Printf("Number of bytes was written: %d", wr)
    
}


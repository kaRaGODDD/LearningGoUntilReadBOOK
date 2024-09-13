/*

Найдите веб-сайт, который содержит большое количество дан
ных. Исследуйте работу кеширования путем двукратного запуска f e t c h a l l и срав
нения времени запросов. Получаете ли вы каждый раз одно и то же содержимое?
Измените f e t c h a l l так, чтобы вывод осуществлялся в файл и чтобы затем можно
было его изучить.


*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch_all(url string, ch chan <- string) {
   
    start := time.Now()

    resp, err := http.Get(url)

    if err != nil {
        ch <- fmt.Sprint("Error was occured while trying contact to server")
    }

    bytes, err := io.Copy(io.Discard, resp.Body)

    if err != nil {
        ch <- fmt.Sprint("Error while copying the result from response")
    }

    resp.Body.Close()
    
    byteCount := fmt.Sprintf("URL: %s, Bytes: %d\n", url, bytes)
    file, err := os.OpenFile("bytes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    _, err = file.WriteString(byteCount)
    defer file.Close()

    end := time.Since(start).Seconds()
    ch <- fmt.Sprintf("Time %.2fs and amoung of bytes %d", end, bytes)

}

func main() {
   
    args := os.Args[1:]
    
    ch := make(chan string) 


    for _, url := range args {
        go fetch_all(url, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
}

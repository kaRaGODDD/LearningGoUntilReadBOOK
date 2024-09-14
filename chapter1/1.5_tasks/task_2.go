package main

import (
    "fmt"
    "os"
    "strings"
    "net/http"
)

func get_resp(url string) {
    _, err := http.Get(url) 
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err)
    }
    
    fmt.Printf("Get the answer from server")


}

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "https://") {

            n_url := "https://" + url
            get_resp(n_url)

   
        } else {
            fmt.Printf("There is http prefix")
            get_resp(url) 
            fmt.Printf("Get the answer from server")

        }
    }
}

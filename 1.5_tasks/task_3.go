package main

import (
	"fmt"
	"os"
    "net/http"
)

func main() {

    args := os.Args[1:]
    
    for _, url := range args {
        
        resp, _ := http.Get(url)
        
        fmt.Printf("All good %s", resp.Status) 


    }
    

}

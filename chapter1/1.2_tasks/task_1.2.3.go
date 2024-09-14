
/*
   :Поэксперементируйте с изменением разницы времени выполнения потенциально
   :неэффективных версий и версии с примененеие strings.Join

*/

package main

import (
	"fmt"
	"time"
	"os"
	"strings"
)

func myself_join_test(args []string, separator* string) {

    start := time.Now()
    
    res := "" 

    for _, value := range args {
       res += value + *separator 
    }
    
    fmt.Println(res)

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func stdlib_join_test(args []string, separator* string) {

    start := time.Now()
    
    res := strings.Join(args, *separator) 
    fmt.Println(res)

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())


}

func main() {
    
    sep := "/"
    
    args := strings.Split(os.Args[0], sep)

    myself_join_test(args, &sep)
    stdlib_join_test(args, &sep)

}

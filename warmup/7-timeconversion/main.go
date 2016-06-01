package main

import (
    "fmt"
    "time"
)

func main() {
    var a string
    
    fmt.Scanf("%v\n", &a)
    
    pt,err := time.Parse("03:04:05PM", a)

    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(pt.Format("15:04:05"))
}
package main

import (
    "fmt"
)

func main() {
    var a, temp, res int
    b :=[]int{}
    
    fmt.Scanf("%v\n", &a)
    
    for i:=0; i < a; i++ {
        fmt.Scanf("%v", &temp)
        b = append(b, temp)
    }    
    
    for i := range b {
        res += b[i]
    }
    
    fmt.Println(res)

}
package main

import (
    "fmt"
)

func main() {
    var temp, res int64
    var a int
    b :=[]int64{}
    
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
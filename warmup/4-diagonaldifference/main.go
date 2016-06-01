package main

import (
    "fmt"
)

func main() {
    var a, temp, diff int
    
    fmt.Scanf("%v\n", &a)
    b := make([][]int, a)
    
    for j := 0; j < a; j++ {
        for i := 0; i < a; i++ {
            fmt.Scanf("%v", &temp)
            b[j] = append(b[j], temp)
        }
        //fmt.Scanf("\n")
    }    
   
    for i := 0; i < a; i++ {
        diff += b[i][i]
    } 
    
    for i, j:=(a-1), 0; i >= 0; i,j = i-1, j+1 {
        diff -= b[j][i]
    }
    
    if diff < 0 {
        diff = -diff
    }
    
    fmt.Println(diff)
}
package main

import (
    "fmt"
)

//TODO reduce big O
func main() {
    var a int
        
    fmt.Scanf("%v\n", &a)
    res := make([][]string, a)
    
    for i:= 0; i < a; i++ {
        res[i] = make([]string, a)
        for j:= a-1; j >=0; j-- {
            if j >= ((a-1) - i) {
                res[i][j] = "#"
            } else {
                res[i][j] = " "
            }
        }
    }

    for i:= range res {
        for j:= range res[i] {
            fmt.Printf("%s", res[i][j])    
        }
        fmt.Printf("\n")
    }
}
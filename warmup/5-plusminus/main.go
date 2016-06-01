package main

import (
    "fmt"
)

func main() {
    var temp int
    var a int

    b :=[]int{}
    res :=[3]int{}
    frqs := [3]float64{}

    fmt.Scanf("%v\n", &a)
    
    for i:=0; i < a; i++ {
        fmt.Scanf("%v", &temp)
        b = append(b, temp)
    }    
    
    lenb := len(b)

    for i:= range b {
        if (b[i] > 0) {
            res[0]++
        }
        
        if b[i] < 0 {
            res[1]++
        }
        
        if b[i] == 0 {
            res[2]++
        }
    }
    
    for i:= range res {
       frqs[i] = float64(res[i])/float64(lenb)
    }
    
    
    for i:= range frqs {
        fmt.Printf("%.4f\n", frqs[i])
    }
    
}
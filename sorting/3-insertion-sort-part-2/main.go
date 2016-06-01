package main

import "fmt"

func main() {
    var shifts, a, j, t int
    b :=[]int{}
    
    fmt.Scanf("%v\n", &a)
    
    for i:=0; i < a; i++ {
        fmt.Scanf("%v", &t)
        b = append(b, t)
    }
    
    for i := 1; i < len(b); i++ {
        t = b[i]
        j = i - 1
        
        for ;j >= 0 && b[j] > t;{
            shifts++
            b[j+1] = b[j]
            j = j - 1
        }
        b[j + 1] = t
    }
    fmt.Println(shifts)
}


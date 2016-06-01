package main

import "fmt"

func main() {
    var a, j, t int
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
            b[j+1] = b[j]
            print(b)
            j = j - 1
        }
        b[j + 1] = t
    }
    print(b)
}

func print(b []int) {
    for j := range b {
        fmt.Printf("%d ", b[j])
    }
    fmt.Printf("\n")
}

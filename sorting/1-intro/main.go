package main
import "fmt"

func main() {
    var v, a, temp int
    b :=[]int{}
    
    fmt.Scanf("%v\n", &v)
    fmt.Scanf("%v\n", &a)
    
    for i:=0; i < a; i++ {
        fmt.Scanf("%v", &temp)
        b = append(b, temp)
    }
    
    for i := range b {
        if (v == b[i]) {
            fmt.Println(i)
            break;
        }
    }
    
}
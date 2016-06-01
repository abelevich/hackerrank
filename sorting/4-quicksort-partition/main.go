package main
import "fmt"


func main() {
    var a, t int
    b :=[]int{}
    
    fmt.Scanf("%v\n", &a)
    
    for i:=0; i < a; i++ {
        fmt.Scanf("%v", &t)
        b = append(b, t)
    }
    
    arr := partion(b)
    print(arr)
}

func partion(b []int ) []int {
    A := []int{}
    eq := []int{}
    r := []int{}
    l := []int{}

	pivot := b[0]
    
    for i := range b {
        if b[i] == pivot {
            eq = append(eq, b[i])
        }
        
        if b[i] < pivot {
            l = append(l, b[i])
        }
        
        if b[i] > pivot {
            r = append(r, b[i])
        }
    }
	
    A = append(A, l...)
    A = append(A, eq...)
    A = append(A, r...)
    
    return A
}


func print(b []int) {
    for j := range b {
        fmt.Printf("%d ", b[j])
    }
    fmt.Printf("\n")
}
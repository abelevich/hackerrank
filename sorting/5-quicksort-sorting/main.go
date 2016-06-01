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
    
    quicksort(&b, 0, len(b))
}

func quicksort(arr *[]int, lo int, hi int) {
	if lo < hi {
        p := partion(arr, lo, hi)
        quicksort(arr, lo, p)
        quicksort(arr, p + 1, hi)
        
        if (hi - lo) > 1 {
            print(*arr, lo, hi)
        }
	}
}

func partion(b *[]int, lo int, hi int ) int {
    r := []int{}
    l := []int{}
    arr := *b
    
    p := lo
    pivot := arr[lo]

    for i := lo + 1; i < hi; i++ {
        if arr[i] > pivot {
            r = append(r, arr[i])
        } else {
            l = append(l, arr[i])
        }
    }
 	
    for i:= 0; i < len(l); i++ {
        arr[p] = l[i]
        p++
    }

    ppos := len(l) + lo
    arr[ppos] = pivot
    
    for i:= 0; i < len(r); i++ {
        ppos++
        arr[ppos] = r[i]
    }

    return lo + len(l) 
}

func print(b []int, lo, hi int) {
    for i:=lo; i < hi; i++ {
        fmt.Printf("%d ", b[i])
    }
    fmt.Printf("\n")
}

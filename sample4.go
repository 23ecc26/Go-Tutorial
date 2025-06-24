package main
import "fmt"

func main() {
    arr := []int{4, 2, 9, 7, 5}
    max := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    fmt.Println(max)
}

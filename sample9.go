package main
import "fmt"

func main() {
    mat1 := [2][2]int{{1, 2}, {3, 4}}
    mat2 := [2][2]int{{5, 6}, {7, 8}}
    var sum [2][2]int

    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            sum[i][j] = mat1[i][j] + mat2[i][j]
        }
    }

    fmt.Println("Sum of matrices:")
    for _, row := range sum {
        fmt.Println(row)
    }
}

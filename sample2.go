package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)

    sum := 0
    for num > 0 {
        sum += num % 10
        num /= 10
    }
    fmt.Println("Sum of digits:", sum)
}

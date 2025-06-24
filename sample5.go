package main
import (
    "fmt"
)

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&str)

    freq := make(map[rune]int)
    for _, ch := range str {
        freq[ch]++
    }

    for k, v := range freq {
        fmt.Printf("%c: %d\n", k, v)
    }
}

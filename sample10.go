package main
import (
    "fmt"
)

func isPalindrome(s string) bool {
    r := []rune(s)
    n := len(r)
    for i := 0; i < n/2; i++ {
        if r[i] != r[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)

    if isPalindrome(input) {
        fmt.Println("Palindrome")
    } else {
        fmt.Println("Not Palindrome")
    }
}

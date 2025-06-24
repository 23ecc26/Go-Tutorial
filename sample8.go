package main
import "fmt"

type Student struct {
    name  string
    age   int
    marks int
}

func main() {
    var s Student
    fmt.Print("Enter name, age, marks: ")
    fmt.Scanln(&s.name, &s.age, &s.marks)

    fmt.Println("Student Info:")
    fmt.Println("Name:", s.name)
    fmt.Println("Age:", s.age)
    fmt.Println("Marks:", s.marks)
}

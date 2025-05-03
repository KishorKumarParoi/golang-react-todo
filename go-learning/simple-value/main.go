package main

import "fmt"

func main() {
    // Integer
    fmt.Println(1)

    // String
    fmt.Println("Hello")

    // Float
    fmt.Println(3.14)

    // Boolean
    fmt.Println(true)
    fmt.Println(false)

    // Complex numbers
    fmt.Println(1 + 2i)

    // Rune (Unicode code point)
    fmt.Println('A')

    // Array
    arr := [3]int{1, 2, 3}
    fmt.Println(arr)

    // Slice
    slice := []string{"Go", "is", "fun"}
    fmt.Println(slice)

    // Map
    m := map[string]int{"one": 1, "two": 2}
    fmt.Println(m)

    // Struct
    type Person struct {
        Name string
        Age  int
    }
    p := Person{Name: "John", Age: 30}
    fmt.Println(p)

    // Pointer
    x := 42
    ptr := &x
    fmt.Println(ptr, *ptr)

    // Function
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(2, 3))

    // Channel
    ch := make(chan int)
    go func() {
        ch <- 10
    }()
    fmt.Println(<-ch)

    // Interface
    var i interface{} = "interface example"
    fmt.Println(i)

    // Nil
    var ptrNil *int
    fmt.Println(ptrNil)
}
package main

import "fmt"

func main() {
    // 1. Declaring variables with `var`
    var a int = 10 // Explicit type
    var b = 20     // Type inferred
    var c string   // Default value (zero value) is an empty string
    fmt.Println("a:", a, "b:", b, "c:", c)

    // 2. Short variable declaration (only inside functions)
    d := 30 // Type inferred
    e := "Hello, Go!"
    fmt.Println("d:", d, "e:", e)

    // 3. Multiple variable declaration
    var x, y, z int = 1, 2, 3
    fmt.Println("x:", x, "y:", y, "z:", z)

    // 4. Grouped variable declaration
    var (
        name  string = "John"
        age   int    = 25
        email string = "john@example.com"
    )
    fmt.Println("name:", name, "age:", age, "email:", email)

    // 5. Constants
    const pi float64 = 3.14159
    const greeting = "Welcome to Go!"
    fmt.Println("pi:", pi, "greeting:", greeting)

    // 6. Zero values
    var zeroInt int
    var zeroFloat float64
    var zeroBool bool
    var zeroString string
    fmt.Println("Zero values -> int:", zeroInt, "float:", zeroFloat, "bool:", zeroBool, "string:", zeroString)

    // 7. Pointers
    var num int = 42
    var ptr *int = &num // Pointer to `num`
    fmt.Println("Pointer -> Address:", ptr, "Value:", *ptr)

    // 8. Type conversion
    var f float64 = 3.5
    var i int = int(f) // Explicit conversion
    fmt.Println("Type conversion -> float:", f, "int:", i)

    // 9. Blank identifier `_`
    _, value := 10, 20 // Ignore the first value
    fmt.Println("Blank identifier -> value:", value)

    // 10. Variables with different types
    var booleanVar bool = true
    var stringVar string = "Go is awesome!"
    var floatVar float64 = 9.81
    fmt.Println("booleanVar:", booleanVar, "stringVar:", stringVar, "floatVar:", floatVar)
}
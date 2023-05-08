# Section 2: Overview of the Go Language

## Variables & Functions

- in Go, we use `var` keyword to declare a variable with Camel case name convention. the syntax is:

    ```go
    var variableName variableType
    ```
- we can assign value to a variable by using `=` operator. the syntax is:

    ```go
    variableName = value
    ```

- then we make use of variable by using `fmt.Println()` function. the syntax is:

    ```go
    fmt.Println(variableName)
    ```

- now let run our program to see the result:

    ```go
    package main

    import "fmt"

    func main() {
        var name string
        name = "John"
        fmt.Println(name)
    }
    ```

    ```shell
    $ go run main.go
    John
    ```
  

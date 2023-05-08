# Section 2: Overview of the Go Language

## Variables & Functions

- in Go, we use `var` keyword to declare a variable with Camel case name convention. the syntax is:

    ```go
    var variableName variableType
    ```
- we must use the variable after we declare it, otherwise we will get an error.
- we can assign value to a variable by using `=` operator. the syntax is:

    ```go
    variableName = value
    ```

- then we make use of variable by using `fmt.Println()` function. the syntax is:

    ```go
    fmt.Println(variableName)
    ```

- checkout the code in [main.go](main.go) file.

- now we create a function that return a string in the `main.go` file.
- we can also declare package level variable by using `var` keyword outside the function.
- function in Go can return more than 1 value.
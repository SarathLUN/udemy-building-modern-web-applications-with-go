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

- we must use the variable after we declare it, otherwise we will get an error.
- here we make use of variable by print it out with `fmt.Println()` function. the syntax is:

  ```go
  fmt.Println(variableName)
  ```

- we can also declare package level variable by using `var` keyword outside any function.
- now we create a function `saySomething` that return a string in the `main.go` file.
- we can use short declaration operator `:=` to declare and assign value to a variable. the syntax is:

  ```go
  variableName := value // the data type will be inferred from the type of value
  ```

- function in Go can return more than 1 value, checkout function `saySomethingElse` that return `string` and `int`.
- checkout the code in [variable-and-function.go](variable-and-function.go) file.

## Pointers

- we start by creating new file: `pointer.go`
- this file will have function `main` & `changeUsingPointer`
- in function `main` we declare a local variable which is not accessible from another function
- however we will use pointer of that variable as param for function `changeUsingPointer` and rewrite new value on it.
- also notice that function `changeUsingPointer` is not return anything.
- if we try to print out the value of pointer it will show the memory address in hexadecimal.
- code of this topic: [pointer.go](pointer.go)

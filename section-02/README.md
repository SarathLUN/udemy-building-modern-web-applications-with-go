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

## Types & Structs

- we start by creating new file: `types-and-structs.go`
- be careful with variable shadowing, when we initialize a new variable, make sure it is unique name.

```go
package main

import "log"

s := "Hello World!"

func main() {
	s := "Something Else!" // variable shadowing 
	log.Println(s)
	// output: Something Else!
}
```

- if we want the struct, field, variable is exposed to outside the package, we need to use capitalize.

Ex:

```go
// other/hello.go
package other

var SomeText = "Hello World!"

```

```go
// main.go
package main

import (
	"log"
	"other"
)

func main() {
	log.Println(other.SomeText)
}
```

- each data type has it own default value:
    - default value of `string` is ""
    - default value of `int` is 0
    - default value of `bool` is false
    - and more [here](https://golangbyexample.com/go-default-zero-value-all-types/)
- when we initialize a variable or a struct without specify it value, the default value will be assigned.

## Receiver

- create file `receiver.go`.
- execute by `go run receiver.go`.
- this is similar to method in OOP concept.
- the benefit of this is we can perform some logic before returning data of the struct.

## Map & Slide

Key takeaways:

1. Data structures are important for storing and organizing data in programs.
2. Simple variables can store individual values of different types.
3. Maps are a data structure used to store key-value pairs. They are created using the "make" keyword and can store
   various data types.
4. Maps are accessed using keys and can be overwritten to update values.
5. Maps are useful for fast data access and are immutable, meaning they stay constant regardless of their usage in
   different parts of the program.
6. Maps are not sorted, so the order of elements cannot be assumed.
7. Slices are another data structure similar to arrays, commonly used in Go.
8. Slices can store multiple values of the same type and can be dynamically expanded using the "append" function.
9. Slices can be sorted and manipulated using various built-in functions.
10. Slices can be declared using the shorthand syntax and can be accessed using index ranges.

- create file `map-slide.go`
- execute by `go run map-slide.go`

### Re-Structure Package `section-02`

- up-to this point, I would like to re-structure the package so that we can run from our IDE.
- the below structure make `main.go` as our entry point and in the function `main` we can call to other package with
  function `Start()`.

```shell
.
|-- README.md
|-- go.mod
|-- main.go
|-- my_map_and_slice
|   `-- map-slice.go
|-- my_pointer
|   `-- pointer.go
|-- my_receiver
|   `-- receiver.go
|-- my_type_and_struct
|   `-- types-and-structs.go
`-- my_variable_and_function
    `-- variable-and-function.go
```

```go
package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/my_map_and_slice"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/my_pointer"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/my_receiver"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/my_type_and_struct"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-02/my_variable_and_function"
)

func main() {
	my_variable_and_function.Start()
	my_pointer.Start()
	my_type_and_struct.Start()
	my_map_and_slice.Start()
	my_receiver.Start()
}

```

## decision

- file `my_decision_structure/decision.go`.
- run the package from function `main` in the `main.go` file.
- we can use `&&` for AND operation and `||` for OR operation.
- `!` is used to check opposite, means `not` 
- syntax `if` statement:

```go
if condition {
action
} else if condition {
action
} else if ... {
action
}

...

} else {
action
}
```

- The switch statement is another decision-making structure that simplifies handling multiple cases. It checks the value
  of a variable against different cases and executes the corresponding code block for the matching case.
- In Go, the switch statement stops executing as soon as it finds a matching case, unlike some other programming
  languages where it continues to evaluate subsequent cases.
- syntax:

```go
switch myVar
case 1:
    action
case 2:
    action
    
...
	
default:
	action
```

## Loop

- file `my_loop/loop-and-range.go`
- execute by `go run my_loop/loop-and-range.go`
- syntax:

```go
for initialization; condition; post {
// Code to be executed repeatedly
}

```

- in this session we also introduce **blank identifier** `_` when we want to ignore the returned value from a function.
- `string` is actually immutable, so when you override the value of a string to something else, it's actually destroying one object in memory and creating an entire new one. To demonstrate this let show the memory address.
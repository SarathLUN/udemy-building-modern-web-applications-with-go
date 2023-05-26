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

## decision

- file `decision.go`.
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

- file `loop-and-range.go`
- execute by `go run loop-and-range.go`
- syntax:

```go
for initialization; condition; post {
// Code to be executed repeatedly
}

```

- in this session we also introduce **blank identifier** `_` when we want to ignore the returned value from a function.
- `string` is actually immutable, so when you override the value of a string to something else, it's actually destroying one object in memory and creating an entire new one. To demonstrate this let show the memory address.

## Interface

- we want to have a function that can operate on 2 or more difference types in param
- then we can create an interface which will stand in the middle between the 2 types and the function that we need to use
- in Go, the 2 types must be certified to be the member of the interface
- to be certified to an interface, those 2 types need to implement the all methods required by the interface
- in example file [`interface.go`](interface.go) we have 2 types: `Dog` and `Gorilla`
- we want to use only 1 function `PrintInfo()` without create method for each type
- so we use interface `Animal` in the middle
- so function `PrintInfo()` take param `Animal` but not `Dog` ether `Gorilla`
- then to be able to past type `Dog` or `Gorilla` into function `PrintInfo()`, they need to be part of `Animal`
- to do so `Dog` and `Gorilla` need to implement required method `Says()` & `NumberOfLegs()` for interface `Animal`
- after certified, `Dog` or `Gorilla` are `Animal` which mean they are acceptable by function `PrintInfo()`
- so we don't need to create separate function for each type.
- however, in Go, most receiver are pointer type as best practice, for faster performance.
- indeed, here are the technical reasons on Why we need to use interface in Go:
  - Abstraction: Interfaces allow you to define abstract types that specify a set of methods without specifying the implementation details. This allows you to work with objects based on their behavior rather than their concrete types. By defining interfaces, you can create code that is more flexible and reusable.
  - Polymorphism: Interfaces enable polymorphism in Go. You can have multiple types that implement the same interface, and you can work with them using the interface type. This allows you to write code that is more generic and can handle different types interchangeably, as long as they satisfy the interface contract.
  - Dependency injection: Interfaces are often used in dependency injection scenarios. By defining interfaces for dependencies, you can write code that depends on those interfaces rather than concrete implementations. This promotes loose coupling between components and makes it easier to swap out implementations without affecting the code that uses them.
  - Testability: Interfaces make it easier to write unit tests for your code. By depending on interfaces instead of concrete types, you can easily create mock implementations of those interfaces for testing purposes. This allows you to isolate the code under test and provide controlled behavior for testing different scenarios.
  - Code extensibility: Interfaces provide a contract that specifies the required behavior. If you want to extend your code by adding new types that satisfy an existing interface, you can do so without modifying the existing code. This makes it easier to extend and evolve your codebase over time.

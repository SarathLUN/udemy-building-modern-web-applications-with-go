# udemy building modern web applications with go

This repository content all my practice, tips, and techniques while learning Udemy
course: [Building Modern Web Applications with Go (Golang)](https://www.udemy.com/course/building-modern-web-applications-with-go/)

I make each section as a branch, so you can easily navigate to the branch you want to see.

Enjoy learning!

## Section 1: Introduction

### What we will do?

- Learn the fundamentals of Go programming language
- Learn how to build a web application with Go
- Learn how to deploy a web application to a cloud server

### why go?

- compiled language, single binary file (fast)
- no runtime to worry about
- statically typed, so no surprises at runtime
- object-oriented (sort of), but not class based
- concurrency built in
- garbage collected
- cross-platform
- excellent package management and testing built in

now, let look at examples:

- the left terminal is running a Go program, and the right terminal is running a PHP program. then we run command `top`
  and `uptime` to see load of each server:

![load](./images/section-01-001.png)

- as we can see that the server that running Go program has a lower load than the server that running PHP program. this
  is because Go is a compiled language, so it's faster than PHP, which is an interpreted language.

### install Go:

- [download Go](https://go.dev/dl/) and install Go for your OS
- start your terminal and run command `go version` to check if Go is installed successfully
- [set up Go workspace](https://golang.org/doc/code.html#Workspaces)
- install your favorite IDE or text editor.
- now we can start coding a small program to test if everything is working fine. create a file `main.go` and add the
  following code:
- first we need to initialize a Go module by running command `go mod init <full package name>`

```shell
$ go mod init github.com/SarathLUN/udemy-building-modern-web-applications-with-go
go: creating new go.mod: module github.com/SarathLUN/udemy-building-modern-web-applications-with-go

```

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

- run command `go run main.go` to run the program. you should see the output `Hello, World!` in your terminal.

```shell
$ go run main.go
Hello, World!

```

## Section 2: Overview of the Go Language

### Variables & Functions

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
  

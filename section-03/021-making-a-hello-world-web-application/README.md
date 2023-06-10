# Section 3: Building a Basic Web Application

## video 21

- in this video we create a simple **Hello World** web service with standard library `net/http`
- function `HandleFunc()` take 2 arguments `pattern` as string and handler as function
- handler function required 2 arguments **ResponseWriter** and pointer of request **\*Request**
- in this video we use closure function to handle the client web request and response with words `Hello World` to the client browser
- function `Fprintf` will write the words **Hello World** to the response and return the number of bytes and an error, so we just log it to console
- finally, we use function `ListenAndServe()` to make the server listen on port **8080**
- next we can run this file by `go run 21-making-a-hello-world-web-application.go`
- then we can access our web server by [`http://localhost:8080`](http://localhost:8080)
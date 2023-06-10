# Section 3: Building a Basic Web Application

## video 24:

- this lesson we create to routes pattern `/` for home page and `/about` for about page
- we also create handler function and not using closure function anymore
- both function `Home()` and `About` is just to render respective page for responding to http client
- we also create **constant** `portNumber` which mean it can't be changed in the running time
- we learn how we handle error 
- we create a handler function `Divide()` to handle request for `/divide`
- function `Divide()` will call to another function `divideValues` that return a float and a potential error
- if we got an error we just write it to the response, so the client web browser can view the error for this exercise.
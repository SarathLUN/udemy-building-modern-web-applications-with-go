# Section 4: Improve Routing & Middleware

## video #038 Using package pat for routing

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/038-Using-package-pat-for-routing
```

- in this session we use package pat to setup our routers.
- we create separate file `cmd/web/router.go`
- this is the github link to package [pat](https://github.com/bmizerany/pat)
- install pat:

```shell
$ go get github.com/bmizerany/pat
```

- import pat, then we create and return multiplexer (mux), in between is our setup routes
- we also need to cast our handler to be HandlerFunc
- in main function we instantiate **http.Server**, and start the server
- 

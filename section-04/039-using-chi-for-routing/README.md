# Section 4: Improve Routing & Middleware

## video #039 Using package chi for routing

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/039-using-chi-for-routing
```

- in this session we use package `chi` to set up our routers.
- what we need to change is file `cmd/web/routers.go`
- we create new router from `chi` instead of `pat` in very easy swapping because we already separate the concern.
- we also introduce middleware `Logger`

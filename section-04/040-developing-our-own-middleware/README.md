# Section 4: Improve Routing & Middleware

## video #040 Developing our own middleware

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/040-developing-our-own-middleware
```

- in this lesson, we create middleware file `cmd/web/middleware.go`
- we create function `NoSurf()` to set base cookie for **CSRF token**
- 
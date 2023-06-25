# Section 3: Building a Basic Web Application

## video 034-Why-the-application-wide-config-is-so-useful

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/034-Why-the-application-wide-config-is-so-useful
```

- we have package `config` to avoid import cycle from each other and your app will not compile
- we already created a struct `AppConfig` and access it in the main function
- in this section, we add more fields `UseCache` as `bool` and `InfoLog` as pointer to `log.Logger`
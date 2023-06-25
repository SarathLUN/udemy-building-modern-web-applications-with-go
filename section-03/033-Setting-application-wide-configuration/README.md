# Section 3: Building a Basic Web Application

## video 033-Setting-application-wide-configuration

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/033-Setting-application-wide-configuration
```

- we create package `config` to avoid import cycle from each other and your app will not compile
- we create a struct `AppConfig` and access it in the main function
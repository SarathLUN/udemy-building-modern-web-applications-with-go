# Section 3: Building a Basic Web Application

## video 035-Optimizing-our-template-cache-by-using-an-application-config

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config
```

- we have package `config` to avoid import cycle from each other and your app will not compile
- we already created a struct `AppConfig` and access it in the main function
- we added more fields `UseCache` as `bool` and `InfoLog` as pointer to `log.Logger`
- we create function `NewTemplate()` to instantiate `AppConfig`
- then in our function `main()` we create new template
- now in the function `RenderTemplate()` we not render template directly, instead we use `app.TemplateCache`
- we can run our server and tested, it works fine.
- now we make use of AppConfig in our handler.
- we use **Repository Pattern**
- we create a struct `Repository` and variable `Repo` which type of Repository.
- we create function `NewRepo()` to instantiate Repository.
- we create function `NewHandlers()` to set repository for handler.
- then we set receiver in our handlers `Home()` and `About()`.
- so in function `main()` when we access to handler need to specify receiver, so it become `handlers.Repo.Home` something like that.
- to demonstrate how it works, let pull something out of AppConfig, in this case let use `UseCache`, then in **main** set the `app.UseCache = true`.
- in the current state, if we change template, it will not load to the client since it read from cache.
- so in function `RenderTemplate()` we check if we want to use existing cache or create new one.
- this time we try to set `app.UseCache = false` so it will always read and render template from disk, not from cache. this is useful for development mode.
- this is just one of example, in the future use `AppConfig` to configure our application behaviour.

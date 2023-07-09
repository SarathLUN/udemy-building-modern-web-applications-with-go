# Section 5: State Management with Sessions

## Lesson #042: Experimenting with session

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/042-experimenting-with-session
```

- this time is the actual use of `Repository` so make sure we initiate it in `main` function.

```go
	// create new repo
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// render new template
	renderers.NewTemplate(&app)
```

- in this lesson, we grep the client IP address everytime user visit our home page.
- then we show back the client IP address in the page `/about`.
- we use function `Put()` to set value into session cookie in `Home()` handler.
- we use function `GetString()` in `About()` handler to query back the client IP address and display it to the web page.
- 
# Section 5: State Management with Sessions

## Lesson #041: Installing and Setting up a sessions package

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/041-install-and-setting-up-a-session-package
```

- we use [SCS: HTTP Session Management for Go](https://github.com/alexedwards/scs)
- we instantiate session manager and set lifetime for 24h.
- this time we use cookie to store the session, so we set the `Persist`, `SameSite`, `Secure`.
- we make use of the `AppConfig` to session available to whole application.
- in the middleware we create new function to actually store the session into cookie.
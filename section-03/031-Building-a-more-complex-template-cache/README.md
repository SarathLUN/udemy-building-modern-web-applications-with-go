# Section 3: Building a Basic Web Application

## video 031-Building-a-more-complex-template-cache

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/031-Building-a-more-complex-template-cache
```

- in package `renderers` we re-create function `createTemplateCache()`
- we find all files with the extension `.page.tmpl`.
- then loop range over the slice `pages`
- each iteration get the file name.
- then we look for layout file which are end with `.layout.tmpl`
- we check if the layout files needed.
- now we can test our server

```shell
go run ./cmd/web
```
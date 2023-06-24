# Section 3: Building a Basic Web Application

## video 030-Building-a-simple-template-cache

- update `go.mod` for new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/030-Building-a-simple-template-cache
```

- in this lesson we create cache of template instead of reading from disk everytime of request.
- the best data structure to use in this case is the map at package level.
- we create a map `tc` to store the template cache.
- in package `renderers` function `RenderTemplate()` we check if the template already exist in the template cache (tc).
- if the template that we are looking for not yet exist in the cache, we render new one by calling to function `createTemplateCache()` then store into template cache.
- or else, we just re-use the existing template from cache.
- 
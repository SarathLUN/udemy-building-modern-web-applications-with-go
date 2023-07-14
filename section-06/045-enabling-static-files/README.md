# Section 6: Choosing a project and working with forms

## Lesson #045: Enabling static files

- copy project from the previous lesson #044 to #045

```shell
cd section-06
cp -rfv 045-enabling-static-files 045-enabling-static-files
cd 045-enabling-static-files
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/045-enabling-static-files # update `go.mod` for new package
goland . # start Goland IDE
```

- then in Goland we can use **Replace in files** (Cmd + Shift + R) to replace "044-setting-up-our-project" by "045-enabling-static-files"
- now in the `home.page.tmpl` we try to load image from `/static/images/house.jpeg`, but it can't find the image file
- because of security reason, by default it is block access, Go web server only allow to access resource via exclusively.
- so in our `routers.go`, we use `http.FileServer()` to create file server by providing `http.Dir()` as parameter.
- then we serve static files by using `mux.Handle()` by providing `http.StripPrefix()` as parameter.
- 
# Section 3: Building a Basic Web Application

## video 29:

- our template layout have some repeated header and footer part of HTML.
- we create file `base.layout.tmpl`
- we use directive engine to create layout, we start with `{{ define "name_of_your_layout" }}` and we're closing this with `{{end}}`
- then inside this, we create placeholder with directive `{{ block "name_of_your_block" . }}` with the `.` at for data transfer, then we're closing this with `{{end}}`.
- we also create another 2 block for CSS and JS which will load for specify page if we needed to.
- by this our layout is ready to use.
- now let go to our `home.page.tmpl` and `about.page.tmpl` to use template layout instead.
- we use directive `{{template}}` to load the template.
- then we define the block to fill into placeholder.
- now we can our server, and in this session we use new technique.

```shell
go run ./cmd/web
```

- the server started, but when we try to access the page, we got an error: `error parsing template: html/template:about.page.tmpl:1:11: no such template "base"`
- this because function `template.ParseFiles()` is parsing base on the named file.
- so we need to add the template file also as the 2nd param.
- final call should be:

```go
parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
```

- now we can start the server again, and test
- and it works!
- however, we need to improve efficiency of rendering the template later on, as now it need to read from disk and rendering every time it loading the page.
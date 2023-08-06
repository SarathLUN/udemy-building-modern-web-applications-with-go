# Section 8: Converting our HTML to Go Templates, and creating handlers 

## Lesson #061: Creating handlers for our forms & adding CSRF protection

- make sure we keep version control safe, now let checkout new branch

```shell
git checkout -b 08-061-creating-handlers-for-our-forms-and-adding-csrf-protection
```

- then we copy previous project for this lesson

```shell
cp -rfv 060-converting-our-pages-to-go-templates 061-creating-handlers-for-our-forms-and-adding-csrf-protection
```

- then we need to replace in files in **Goland** press `CMD + SHIFT + R`, replace "061-creating-handlers-for-our-forms-and-adding-csrf-protection" with "061-creating-handlers-for-our-forms-and-adding-csrf-protection"

- we also need to update our go mod file.

```shell
cd section-08/061-creating-handlers-for-our-forms-and-adding-csrf-protection/
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/061-creating-handlers-for-our-forms-and-adding-csrf-protection
```

- now if we run our application, we have few form that doesn't submit to do anything yet. Let start from form **Search for Availability**
- currently, attribute `action=""`, typically we can post to same URL because we will use difference method

```html
<form action="/search-availability" method="post" novalidate class="needs-validation">
```

- as you can see, we will use method `POST` for this route, let create a route to handle this.
- file `routers.go` add one more route that will handle by handler function called `PostAvailability`

```go
mux.Post("/search-availability", handlers.Repo.PostAvailability)
```

- let create the handler function in package `handlers`

```go
// PostAvailability search for availability 
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted to search availability"))
}
```

- for now, we just write to browser a string to see if the post form is working correct.
- let re-start the application and test
- Ooop! we got an error `Bad Request` on the browser instead, what happen?
- well, we had tell the router to use middleware `NoSurf` which is required `CSRF token` to be able to submit a form from client to server.

```go
	mux.Use(NoSurf)
```

- so let generate the token in the function `AddDefaultData` in package `renderers`

```go
td.CSRFToken = nosurf.Token(r)
```

- function `Token()` required a pointer to `*http.Request`, so we need to refactor the function `AddDefaultData` to add one more parameter.
- then we will get another error because function `RenderTemplate` doesn't have `*http.Request` yet, so we also need to refactor this function.

```go
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData)
```

- Goland will auto insert `nil` at the place of `r`

```go
renderers.RenderTemplate(w, nil, "home.page.tmpl", &models.TemplateData{})
```

- we need to replace all with the `r`

```go
renderers.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
```

- the `nosurf` middleware require to have embedded CSRF token when submit to the router, so in form **Search for Availability** add one more input and has to be named as `csrf_token`

```html
<input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
```

- now we are able to submit form, let start to capture the form data
- so our function `PostAvailability` is now should look like below:

```go
// PostAvailability search for availability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	w.Write([]byte(fmt.Sprintf("start date is %s, and end date is %s", start, end)))
}
```

- and when we try to submit we should get text printed on the browser as below:

```shell
start date is 2023-08-07, and end date is 2023-08-09
```
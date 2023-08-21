# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #066: Server-side form validation

- first we need to refactor our project
- replace all path of imported packages from `/065-refactoring-to-use-internal-packages/`  to `/066-server-side-form-validation/`
- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./section-08/066-server-side-form-validation/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/065-refactoring-to-use-internal-packages/#/066-server-side-form-validation/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/066-server-side-form-validation
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- if there is no error and we can access from browser, means it working fine.
- now our clone package is working fine, and in this lesson we will prepare our Go back end package to validate the front end form
- we create one more package with 2 files as below
- with nvim-tree (File Explorer) move the caret to directory `internal` then use `a` command to add new directory and file. So press `a` then type `form/errors.go` then press `Enter`. nvim-tree will create new file for us.
- then move the caret to directory **form** that we just created, and hit `a` again to create another file `forms.go`.
- both files make sure they under package `forms`

```go
package forms

```

- then in the file [`errors.go`](./internal/form/errors.go) we create a maps and a function `Add` to add an error message for a given form field and function `Get` that return error message of a given form field.
- in file [`forms.go`](./internal/form/forms.go) we will create a struct `Form` and function `New` for initialize a form struct.
- let create a function for Form to check required of a given form field call `Has`.
- now let create another route to handle post request from form submit and point to handler function `PostReservation`.
- then in `handlers.go` we create new function for `PostReservation`.
- then we add one more member to our `TemplateData`. So when we rendering the reservation form we need to paste in the form object.
- in form `make-reservation.page.tmpl` don't forget to have input hidden of `csrf_token`
- now we should be able to test submitting the form make reservation and reach the function `PostReservation`.
- And the great thing about this is return error + user input data from back end to front end, so that user no need to re-input again.
- by then, we finished our lesson #66.

See you in the next lesson!


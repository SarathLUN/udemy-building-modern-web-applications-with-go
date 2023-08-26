# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #067: Server-side form validation 2

- so now just copy our project from previous lesson, so we need to refactor our project
- replace all path of imported packages from `/066-server-side-form-validation/`  to `/067-server-side-form-validation-2/`
- first, let update this `README.md` file, using below vim command:

```shell
:%s#/066-server-side-form-validation/#/067-server-side-form-validation-2/#g

```

- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./section-08/067-server-side-form-validation-2/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/066-server-side-form-validation/#/067-server-side-form-validation-2/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/067-server-side-form-validation-2
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- if there is no error and we can access from browser, means it working fine.
- now our clone package is working fine, and in this lesson we will continue our back end validation.
- first, when user submit form to our handler `PostReservation` let parse the form and handle the error if any.
- then, let create new models call `models.go`.
- then, inside this we create a struct called `Reserversion`
- then, back to our handler `PostReservation` and create an object Reserversion
- so now we can make use of our previous method `Has` to validate form fields
- in this lesson, we also introduce new method `Valid` to check the validity of the field 
- also make sure we include error in our existing method `Has` and add the message.
- then we check form valid, if not, we return error message along with input data.
- then in the template file `make-reservation.page.tmpl` we add some logic to display the error that return from back end, in this case it associate with field: `first_name`.
- we use **Go template directives** call `with` to check if any error on field `first_name` then display it out.
- and that's it for this lesson.
- now we run out server and test.

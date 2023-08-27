# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #068: Server-side form validation 3

- so now just copy our project from previous lesson, so we need to refactor our project
- replace all path of imported packages from `/067-server-side-form-validation-2/`  to `/068-server-side-form-validation-3/`
- first, let update this `README.md` file, using below vim command:

```shell
:%s#/067-server-side-form-validation-2/#/068-server-side-form-validation-3/#g

```

- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./section-08/068-server-side-form-validation-3/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/067-server-side-form-validation-2/#/068-server-side-form-validation-3/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/068-server-side-form-validation-3
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- so now we have done refactoring the cloned package, let start our new lesson.
- to be able to apply logic on form loaded, we need to have at least a blank form object on load.
- so the function `Reservation` we need to have blank reservation object.
- along with the errors, we also can return the value from user input in the attribute call `value`.
- we declare variable in golang template to store the reservation object.
- now let restart our application and test by keeping first name blank to get the error, and fill in the other 3 fields and see if the value come back after it reloaded, and it works!
- next, let apply validation and error on all fields, but instead of using method `Has` on all fields, we're creating one more method `Required` which take Variadic function of string so that we can pass many string (in this case field name).
- in that method we range over those fields and apply logic.

```go
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

```

- then go back to our handler function `PostReservation` and not using `Has` to check anymore, but use function `Required` to check instead.

```go
	form.Required("first_name","last_name","email")

```

- now let restart our application and test once again. And it works!
- next let add one more validation method `MinLength` to set minimum char on a field.
- now let test again, and it works!
- for the method `Has` we will use it to validate the checkbox field.
- by then, we are finishing lesson 68. Let continue to next lesson.
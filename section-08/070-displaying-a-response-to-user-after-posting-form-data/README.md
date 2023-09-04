# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #069: Server-side form validation 4

- so now just copy our project from previous lesson, so we need to refactor our project
- replace all path of imported packages from `/068-server-side-form-validation-3/`  to `/069-server-side-form-validation-4/`
- first, let update this `README.md` file, using below vim command:

```shell
:%s#/068-server-side-form-validation-3/#/069-server-side-form-validation-4/#g

```

- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./section-08/069-server-side-form-validation-4/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/068-server-side-form-validation-3/#/069-server-side-form-validation-4/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/069-server-side-form-validation-4
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- so now we have done refactoring the cloned package, let start our new lesson.
- in this lesson we will use 3rd party library to validate field email address call [govalidator](https://github.com/asaskevich/govalidator).
- let install the package

```shell
go get github.com/asaskevich/govalidator
```

- then in file `forms.go` we add one more validation method on form

```go
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
```

- then in our handler function we apply the check on field email.

```go
form.IsEmail("email")
```

- that's it, now we can restart server and test, and it works as expected.
- by then we are ending this lesson 69 here, see you in the next lesson 70.
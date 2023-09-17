# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #071: Finishing up our response to user and add alert

- so now just copy our project from previous lesson, so we need to refactor our project
- replace all path of imported packages from `/070-displaying-a-response-to-user-after-posting-form-data/`  to `/071-finishing-up-our-response-to-user-and-add-alerts/`
- first, let update this `README.md` file, using below vim command:

```shell
:%s#/070-displaying-a-response-to-user-after-posting-form-data/#/071-finishing-up-our-response-to-user-and-add-alerts/#g

```

- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./section-08/071-finishing-up-our-response-to-user-and-add-alerts/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/070-displaying-a-response-to-user-after-posting-form-data/#/071-finishing-up-our-response-to-user-and-add-alerts/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/071-finishing-up-our-response-to-user-and-add-alerts
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- so now we have done refactoring the cloned package, let start our new lesson.

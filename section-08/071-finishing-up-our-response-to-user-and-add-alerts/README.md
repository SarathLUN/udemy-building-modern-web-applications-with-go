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
- currently, if user enter URL (/reservation-summary) directly, they will get blank page, and it's not good UX.
- so we want to redirect them to somewhere else if there is no session with alert box, and the good place for now is `/make-reservation`
- in this lesson we will use template data `FlashMessage, WarningMessage, ErrorMessage`, so let create it in the method `AddDefaultData` in file `render.go`
- now, back to our handler function `ReservationSummary` we update the `if` block to add the error message in the session
- then in base layout template and javascript block, we want to add alert box there via our existing function `notify()`
- we introduce another Golang template directive call `with` which is simply find in template data.
- all done, let test our application, and it's working well.
- best practice of working with session is to destroy it after we're done.
- base in function `ReservationSummary` after getting value from session into variable, let destroy it.
- That's it, we are finishing lesson 71, let continue lesson 72.

keep moving forward, happy learning!
# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #070: Display a response to user after posting form data

- so now just copy our project from previous lesson, so we need to refactor our project
- replace all path of imported packages from `/069-server-side-form-validation-4/`  to `/070-displaying-a-response-to-user-after-posting-form-data/`
- first, let update this `README.md` file, using below vim command:

```shell
:%s#/069-server-side-form-validation-4/#/070-displaying-a-response-to-user-after-posting-form-data/#g

```

- in neovim use below command to add all **go** files into the search scope 

```vim
:args ./section-08/070-displaying-a-response-to-user-after-posting-form-data/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```vim
:argdo %s#/069-server-side-form-validation-4/#/070-displaying-a-response-to-user-after-posting-form-data/#g |update
```

- then edit go mod for our package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/070-displaying-a-response-to-user-after-posting-form-data
```

- now we can start our server and test to make sure it works.

```shell
go run cmd/web/*.go
```

- so now we have done refactoring the cloned package, let start our new lesson.
- in this lesson we need to display the data that submit from `PostReservation` in a summary page.
- let create new golang template called `reservation-summary.page.tmpl` by copying from page `about.page.tmpl`
- this new template should have below code:

```html

{{template "base" .}}

{{define "content"}}
    <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-5">Reservation Summary</h1>
                <hr>

                <table class="table table-striped">
                  <thead></thead>
                  <tbody>
                    <tr>
                      <td>Name:</td>
                      <td>?</td>
                    </tr>
                    <tr>
                      <td>Arrival:</td>
                      <td>?</td>
                    </tr>
                    <tr>
                      <td>Departure:</td>
                      <td>?</td>
                    </tr>
                    <tr>
                      <td>Email:</td>
                      <td>?</td>
                    </tr>
                    <tr>
                      <td>Phone:</td>
                      <td>?</td>
                    </tr>
                  </tbody>
                </table>

            </div>
        </div>
    </div>
{{end}}
```

- next let create a handler function to render this page.
- in our handlers file add one more method called `ReservationSummary`
- then in our router add another route with for `/reservation-summary`
- in our previous handler method `PostReservation` we will set the reservation data into session, then redirect to reservation-summary
- then in the main file we need to register data type into session in our case is `models.Reservation` from standard library `encoding/gob`
- then in function `ReservationSummary` we need to get reservation data from session, for this we need to also check data availability, also make sure we assert the type `models.Reservation` on the found data.
- now in our template we can access data and display it
- that's it, so now we are able to capture form post data and display it on another template.

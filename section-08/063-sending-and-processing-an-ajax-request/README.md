# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #062: Creating a handler that return JSON

- first thing first, let create new git branch to secure our source code with version control

```shell
git checkout -b 08-062-creating-a-handler-that-return-json
```

- then let copy project from previous lesson

```shell
cd section-08
cp -rfv 061-creating-handlers-for-our-forms-and-adding-csrf-protection 062-creating-a-handler-that-return-json
cd 062-creating-a-handler-that-return-json
```

- then we need to update go mod file

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/062-creating-a-handler-that-return-json
```

- start Goland from our new project folder

```shell
goland .
```

- in Goland, replace in files by `CMD + SHIFT + R` and find: "
  /061-creating-handlers-for-our-forms-and-adding-csrf-protection/" replace with "
  /062-creating-a-handler-that-return-json/"
- As I am using Goland, I also need to refactor the project name.
- for this learning purpose we create a new handler function called `AvailabilityJSON` in package handlers

```go
// AvailabilityJSON search for availability and response JSON
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

}
```

- then we add another route, and we use method `get` at this time to be easier test

```go
mux.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)
```

- then we need to create a struct to hold the json data

```go
type jsonResponse struct {
OK      bool   `json:"ok"`
Message string `json:"message"`
}
```

- now let implement our handler to response JSON, it should look like below now

```go
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
resp := jsonResponse{
OK:      true,
Message: "available",
}
out, err := json.Marshal(resp)
if err != nil {
log.Println("error marshal json:", err.Error())
}
w.Write(out)
}
```

- this should work the most, but we want the client browser know that this is the json format, so we add the http header

```go
w.Header().Set("Content-Type", "application/json")
```

- let restart application and test
- at this time we can test in the browser, and we should get json response

```json
{
  "ok": true,
  "message": "available"
}
```
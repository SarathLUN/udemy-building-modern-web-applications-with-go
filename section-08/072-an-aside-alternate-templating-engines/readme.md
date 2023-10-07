# Section 8: Converting our HTML to Go Templates, and creating handlers

## Lesson #072: An Aside: Alternate Templates Engines

- so now just copy our project, we need to refactor our project
- first we need to update the go mod file by remove existing one and initial new one.

```shell
rm -rfv go.mod go.sum
go mod init github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/072-an-aside-alternate-templating-engines
```

- then in file `routes.go` we need to update the import 

```go
import (
	"github.com/bmizerany/pat"
	"net/http"
	"websockets-course/internal/handlers" //this need to change
)
```

- change to:

```go
import (
	"github.com/bmizerany/pat"
	"net/http"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/072-an-aside-alternate-templating-engines/internal/handlers" //changed to this.
)
```

- now, let try to start our application if any error.

```shell
go run cmd/web/*.go
```

- and it's working as expected, it show the error detail if something wrong with our template.
- in this lesson, we demonstrate another template engine called [`Jet`](https://github.com/CloudyKit/jet)
- that's it, we are finishing lesson 72. Let continue to lesson 73.

Keep moving forward, Happy Learning!...


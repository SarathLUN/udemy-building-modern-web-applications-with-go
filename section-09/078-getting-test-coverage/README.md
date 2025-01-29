# Section 9: Writing Tests

## Lesson #078: Getting test coverage

- To get started lesson 78, we need to create new git branch

```shell
git checkout -b 09-078-getting-test-coverage

```

- then we need to continue our code by copying from `077-writing-tests-for-our-Render-package`

```shell
cd section-09
cp -rfv 077-writing-tests-for-our-Render-package 078-getting-test-coverage
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/078-getting-test-coverage
```

### if you use IDE such as VSCode or JetBrains, you can find and replace:
- then we need to find and replace some import path to make it import from correct path
    - find: `section-09/077-writing-tests-for-our-Render-package`
    - replace: `section-09/078-getting-test-coverage`
    - filter: `*.go`

### if you use NeoVIM like me, we need to use command `args` to set the search scope, then use command `argdo` to replace:
- in neovim use below command to add all `*.go` files into the search scope
```shell
:args ./section-09/078-getting-test-coverage/**/*.go
```

- then use `argdo` to find and replace all with `#` as delimiter

```shell
:argdo %s#/077-writing-tests-for-our-Render-package#/078-getting-test-coverage#g |update
```

- now the new package is ready for lesson 78 is ready to follow along the tutorial, let start coding.
- we check our test coverage by using command `go test -cover`

```shell
2025/01/29 13:53:03 {"ok":true,"message":"available"}
2025/01/29 13:53:03 start running PostReservation
2025/01/29 13:53:03 end PostReservation
2025/01/29 13:53:03 cannot get item from session
PASS
coverage: 79.6% of statements
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/078-getting-test-coverage/internal/handlers      0.789s
```

- we can also extract into html version `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

```shell
2025/01/29 14:01:50 {"ok":true,"message":"available"}
2025/01/29 14:01:50 start running PostReservation
2025/01/29 14:01:50 end PostReservation
2025/01/29 14:01:50 cannot get item from session
PASS
coverage: 79.6% of statements
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/078-getting-test-coverage/internal/handlers      0.684s
```


# Section 9: Writing Tests

## Lesson #073: Writing tests for our main package

- To get started lesson 73, we need to create new git branch

```shell
git checkout -b 09-073-writing-tests-for-our-main-package
```

- then we need to continue our code from `071-finishing-up-our-response-to-user-and-add-alerts`

```shell
cp -rfv section-08/071-finishing-up-our-response-to-user-and-add-alerts section-09/073-writing-test-for-our-main-package
```

- then update go module path to our new package

```shell
go mod edit -module github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/073-writing-test-for-our-main-package
```

- then we need to find and replace some import path to make it import from correct path
    - find: `section-08/071-finishing-up-our-response-to-user-and-add-alerts`
    - replace: `section-09/073-writing-test-for-our-main-package`
    - filter: `*.go`
- now the new package for lesson 73 is ready to follow along the tutorial, let start coding.
- after writing `main_test.go`, we can run test in our terminal

```shell
cd cmd/web
go test
```

- output:

```shell
PASS
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/073-writing-test-for-our-main-package/cmd/web    0.316s
```

- we can also use `-v` for output with verbose

```shell
=== RUN   TestRun
--- PASS: TestRun (0.00s)
PASS
ok      github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/073-writing-test-for-our-main-package/cmd/web    0.608s
```

- let also write the test for middleware, create new file in same location: `middleware_test.go` then create function: `TestNoSurf` and `TestSessionLoad`
- now we also want to test our `routers.go` by creating the test file: `routers_test.go`.
- then we can also check our test coverage with command: `go test -cover` this will show in command console.
- we can extract test coverage to html with this command: `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`